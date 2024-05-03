// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate ../../../../../.tools/genqlient

package githubscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/gitproviderreceiver/internal/scraper/githubscraper"

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/gitproviderreceiver/internal/metadata"
)

var errClientNotInitErr = errors.New("http client not initialized")

type githubScraper struct {
	client   *http.Client
	cfg      *Config
	settings component.TelemetrySettings
	logger   *zap.Logger
	mb       *metadata.MetricsBuilder
	rb       *metadata.ResourceBuilder
}

func (ghs *githubScraper) start(ctx context.Context, host component.Host) (err error) {
	ghs.logger.Sugar().Info("starting the GitHub scraper")
	ghs.client, err = ghs.cfg.ToClient(ctx, host, ghs.settings)
	return
}

func newGitHubScraper(
	_ context.Context,
	settings receiver.CreateSettings,
	cfg *Config,
) *githubScraper {
	return &githubScraper{
		cfg:      cfg,
		settings: settings.TelemetrySettings,
		logger:   settings.Logger,
		mb:       metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings),
		rb:       metadata.NewResourceBuilder(cfg.ResourceAttributes),
	}
}

// scrape and return github metrics
func (ghs *githubScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	if ghs.client == nil {
		return pmetric.NewMetrics(), errClientNotInitErr
	}

	now := pcommon.NewTimestampFromTime(time.Now())
	ghs.logger.Sugar().Debug("current time", zap.Time("now", now.AsTime()))

	currentDate := time.Now().Day()
	ghs.logger.Sugar().Debugf("current date: %v", currentDate)

	genClient, restClient, err := ghs.createClients()
	if err != nil {
		ghs.logger.Sugar().Errorf("unable to create clients", zap.Error(err))
	}

	// Do some basic validation to ensure the values provided actually exist in github
	// prior to making queries against that org or user value
	loginType, err := ghs.login(ctx, genClient, ghs.cfg.GitHubOrg)
	if err != nil {
		ghs.logger.Sugar().Errorf("error logging into GitHub via GraphQL", zap.Error(err))
		return ghs.mb.Emit(), err
	}

	// Generate the search query based on the type, org/user name, and the search_query
	// value if provided
	sq := genDefaultSearchQuery(loginType, ghs.cfg.GitHubOrg)

	if ghs.cfg.SearchQuery != "" {
		sq = ghs.cfg.SearchQuery
		ghs.logger.Sugar().Debugf("using search query where query is: %q", ghs.cfg.SearchQuery)
	}

	// Get the repository data based on the search query retrieving a slice of branches
	// and the recording the total count of repositories
	repos, count, err := ghs.getRepos(ctx, genClient, sq)
	if err != nil {
		ghs.logger.Sugar().Errorf("error getting repo data", zap.Error(err))
		return ghs.mb.Emit(), err
	}

	ghs.mb.RecordGitRepositoryCountDataPoint(now, int64(count))

	// Get the branch count (future branch data) for each repo and record the given metrics
	var wg sync.WaitGroup

	for _, repo := range repos {
		repo := repo
		name := repo.Name
		trunk := repo.DefaultBranchRef.Name

		wg.Add(1)
		go func() {
			defer wg.Done()

			branches, count, err := ghs.getBranches(ctx, genClient, name, trunk)
			if err != nil {
				ghs.logger.Sugar().Errorf("error getting branch count for repo %s", zap.Error(err), repo.Name)
			}
			ghs.mb.RecordGitRepositoryBranchCountDataPoint(now, int64(count), name)

			// Iterate through the branches populating the Branch focused
			// metrics
			for _, branch := range branches {
				// Check if the branch is the default branch or if it is not
				// behind the default branch behindby means if the current
				// branch doesn't have any changes, then we'll continue on . we
				// won't be able to get any information that would allow us to
				// actually caculate the branch time with any measure of
				// accuracy. at best we could proceed and only check the
				// aheadby commits, at which might, or might not give us some
				// type of range. this is wasteful at best.
				if branch.Name == branch.Repository.DefaultBranchRef.Name || branch.Compare.BehindBy == 0 {
					continue
				}
				ghs.logger.Sugar().Debugf(
					"default branch behind by: %d\n %s branch behind by: %d in repo: %s",
					branch.Compare.BehindBy, branch.Name, branch.Compare.AheadBy, branch.Repository.Name)

				// Yes, this looks weird. The aheadby metric is referring to the number of commits the branch is AHEAD OF the
				// default branch, which in the context of the query is the behind by value. See the above below comment about
				// BehindBy vs AheadBy.
				ghs.mb.RecordGitRepositoryBranchCommitAheadbyCountDataPoint(now, int64(branch.Compare.BehindBy), branch.Repository.Name, branch.Name)
				ghs.mb.RecordGitRepositoryBranchCommitBehindbyCountDataPoint(now, int64(branch.Compare.AheadBy), branch.Repository.Name, branch.Name)

				var adds int
				var dels int
				var age int64

				adds, dels, age, err = ghs.getCommitInfo(genClient, branch.Repository.Name, branch)
				if err != nil {
					ghs.logger.Sugar().Errorf("error getting commit info: %v", zap.Error(err))
					continue
				}

				ghs.mb.RecordGitRepositoryBranchTimeDataPoint(now, age, branch.Repository.Name, branch.Name)
				ghs.mb.RecordGitRepositoryBranchLineAdditionCountDataPoint(now, int64(adds), branch.Repository.Name, branch.Name)
				ghs.mb.RecordGitRepositoryBranchLineDeletionCountDataPoint(now, int64(dels), branch.Repository.Name, branch.Name)

			}

			// Get the contributor count for each of the repositories
			contribs, err := ghs.getContributorCount(ctx, restClient, name)
			if err != nil {
				ghs.logger.Sugar().Errorf("error getting contributor count for repo %s", zap.Error(err), repo.Name)
			}
			ghs.mb.RecordGitRepositoryContributorCountDataPoint(now, int64(contribs), name)

			// Get Pull Request data
			prs, err := ghs.getPullRequests(ctx, genClient, name)
			if err != nil {
				ghs.logger.Sugar().Errorf("error getting pull requests for repo %s", zap.Error(err), repo.Name)
			}

			var merged int
			var open int

			for _, pr := range prs {
				if pr.Merged {
					merged++

					age := getAge(pr.CreatedAt, pr.MergedAt)

					ghs.mb.RecordGitRepositoryPullRequestTimeToMergeDataPoint(now, age, name, pr.HeadRefName)

				} else {
					open++

					age := getAge(pr.CreatedAt, now.AsTime())

					ghs.mb.RecordGitRepositoryPullRequestTimeOpenDataPoint(now, age, name, pr.HeadRefName)

					if pr.Reviews.TotalCount > 0 {
						age := getAge(pr.CreatedAt, pr.Reviews.Nodes[0].CreatedAt)

						ghs.mb.RecordGitRepositoryPullRequestTimeToApprovalDataPoint(now, age, name, pr.HeadRefName)
					}
				}
			}

			ghs.mb.RecordGitRepositoryPullRequestCountDataPoint(now, int64(open), metadata.AttributePullRequestStateOpen, name)
			ghs.mb.RecordGitRepositoryPullRequestCountDataPoint(now, int64(merged), metadata.AttributePullRequestStateMerged, name)
		}()
	}

	wg.Wait()

	// Set the resource attributes and emit metrics with those resources
	ghs.rb.SetGitVendorName("github")
	ghs.rb.SetOrganizationName(ghs.cfg.GitHubOrg)

	res := ghs.rb.Emit()
	return ghs.mb.Emit(metadata.WithResource(res)), nil
}
