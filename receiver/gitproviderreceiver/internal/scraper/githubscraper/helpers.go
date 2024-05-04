// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package githubscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/gitproviderreceiver/internal/scraper/githubscraper"

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/google/go-github/v61/github"
	"go.uber.org/zap"
)

const (
	// The default public GitHub GraphQL Endpoint
	defaultGraphURL = "https://api.github.com/graphql"
	// The default maximum number of items to be returned in a GraphQL query.
	defaultReturnItems = 100
)

func (ghs *githubScraper) getRepos(
	ctx context.Context,
	client graphql.Client,
	searchQuery string,
) ([]SearchNodeRepository, int, error) {
	// here we use a pointer to a string so that graphql will receive null if the
	// value is not set since the after: $repoCursor is optional to graphql
	var cursor *string
	var repos []SearchNodeRepository
	var count int

	for next := true; next; {
		r, err := getRepoDataBySearch(ctx, client, searchQuery, cursor)
		if err != nil {
			ghs.logger.Sugar().Errorf("error getting repo data", zap.Error(err))
			return nil, 0, err
		}

		for _, repo := range r.Search.Nodes {
			if r, ok := repo.(*SearchNodeRepository); ok {
				repos = append(repos, *r)
			}
		}

		count = r.Search.RepositoryCount
		cursor = &r.Search.PageInfo.EndCursor
		next = r.Search.PageInfo.HasNextPage
	}

	return repos, count, nil
}

func (ghs *githubScraper) getBranches(
	ctx context.Context,
	client graphql.Client,
	repoName string,
	defaultBranch string,
) ([]BranchNode, int, error) {
	var cursor *string
	var count int
	var branches []BranchNode

	for next := true; next; {
		r, err := getBranchData(ctx, client, repoName, ghs.cfg.GitHubOrg, 50, defaultBranch, cursor)
		if err != nil {
			ghs.logger.Sugar().Errorf("error getting branch data", zap.Error(err))
			return nil, 0, err
		}
		count = r.Repository.Refs.TotalCount
		cursor = &r.Repository.Refs.PageInfo.EndCursor
		next = r.Repository.Refs.PageInfo.HasNextPage
		branches = append(branches, r.Repository.Refs.Nodes...)
	}
	return branches, count, nil
}

// Login via the GraphQL checkLogin query in order to ensure that the user
// and it's credentials are valid and return the type of user being authenticated.
func (ghs *githubScraper) login(
	ctx context.Context,
	client graphql.Client,
	owner string,
) (string, error) {
	var loginType string

	// The checkLogin GraphQL query will always return an error. We only return
	// the error if the login response for User and Organization are both nil.
	// This is represented by checking to see if each resp.*.Login resolves to equal the owner.
	resp, err := checkLogin(ctx, client, ghs.cfg.GitHubOrg)

	// These types are used later to generate the default string for the search query
	// and thus must match the convention for user: and org: searches in GitHub
	switch {
	case resp.User.Login == owner:
		loginType = "user"
	case resp.Organization.Login == owner:
		loginType = "org"
	default:
		return "", err
	}

	return loginType, nil
}

// Returns the default search query string based on input of owner type
// and GitHubOrg name with a default of archived:false to ignore archived repos
func genDefaultSearchQuery(ownertype string, ghorg string) string {
	return fmt.Sprintf("%s:%s archived:false", ownertype, ghorg)
}

// Returns the graphql and rest clients for GitHub.
// By default, the graphql client will use the public GitHub API URL as will
// the rest client. If the user has specified an endpoint in the config via the
// inherited ClientConfig, then the both clients will use that endpoint.
// The endpoint defined needs to be the root server.
// See the GitHub documentation for more information.
// https://docs.github.com/en/graphql/guides/forming-calls-with-graphql#the-graphql-endpoint
// https://docs.github.com/en/enterprise-server@3.8/graphql/guides/forming-calls-with-graphql#the-graphql-endpoint
// https://docs.github.com/en/enterprise-server@3.8/rest/guides/getting-started-with-the-rest-api#making-a-request
func (ghs *githubScraper) createClients() (gClient graphql.Client, rClient *github.Client, err error) {
	rClient = github.NewClient(ghs.client)
	gClient = graphql.NewClient(defaultGraphURL, ghs.client)

	if ghs.cfg.ClientConfig.Endpoint != "" {

		// Given endpoint set as `https://myGHEserver.com` we need to join the path
		// with `api/graphql`
		gu, err := url.JoinPath(ghs.cfg.ClientConfig.Endpoint, "api/graphql")
		if err != nil {
			ghs.logger.Sugar().Errorf("error joining graphql endpoint: %v", err)
			return nil, nil, err
		}
		gClient = graphql.NewClient(gu, ghs.client)

		// The rest client needs the endpoint to be the root of the server
		ru := ghs.cfg.ClientConfig.Endpoint
		rClient, err = github.NewClient(ghs.client).WithEnterpriseURLs(ru, ru)
		if err != nil {
			ghs.logger.Sugar().Errorf("error creating enterprise client: %v", err)
			return nil, nil, err
		}
	}

	return gClient, rClient, nil
}

// Get the contributor count for a repository via the REST API
func (ghs *githubScraper) getContributorCount(
	ctx context.Context,
	client *github.Client,
	repoName string,
) (int, error) {
	var all []*github.Contributor

	// Options for Pagination support, default from GitHub was 30
	// https://docs.github.com/en/rest/repos/repos#list-repository-contributors
	opt := &github.ListContributorsOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	for {
		contribs, resp, err := client.Repositories.ListContributors(ctx, ghs.cfg.GitHubOrg, repoName, opt)
		if err != nil {
			ghs.logger.Sugar().Errorf("error getting contributor count", zap.Error(err))
			return 0, err
		}

		all = append(all, contribs...)
		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}

	return len(all), nil
}

// Get the pull request data from the GraphQL API.
func (ghs *githubScraper) getPullRequests(
	ctx context.Context,
	client graphql.Client,
	repoName string,
) ([]PullRequestNode, error) {
	var prCursor *string
	var pullRequests []PullRequestNode

	for hasNextPage := true; hasNextPage; {
		prs, err := getPullRequestData(
			ctx,
			client,
			repoName,
			ghs.cfg.GitHubOrg,
			100,
			prCursor,
			[]PullRequestState{"OPEN", "MERGED"},
		)
		if err != nil {
			return nil, err
		}

		pullRequests = append(pullRequests, prs.Repository.PullRequests.Nodes...)
		prCursor = &prs.Repository.PullRequests.PageInfo.EndCursor
		hasNextPage = prs.Repository.PullRequests.PageInfo.HasNextPage
	}

	return pullRequests, nil
}

func (ghs *githubScraper) evalCommits(
	ctx context.Context,
	client graphql.Client,
	repoName string,
	branch BranchNode,
) (additions int, deletions int, age int64, err error) {
	var cursor *string
	items := 100

	// We're using BehindBy here because we're comparing against the target
	// branch, which is the default branch. In essence the response is saying
	// the default branch is behind the queried branch by X commits which is
	// the number of commits made to the queried branch but not merged into
	// the default branch. Doing it this way involves less queries because
	// we don't have to know the queried branch name ahead of time.
	// We also have to calculate the number of pages because when querying for
	// commits you have to know the exact amount of commits diverged otherwise
	// you'll get all commits from both trunk and the branch from all time.
	pages := getNumPages(float64(defaultReturnItems), float64(branch.Compare.BehindBy))

	for page := 1; page <= pages; page++ {
		if page == pages {
			// On the last page, we need to make sure that the last page is
			// retrieved, so if the remainder is 0 we'll reset to 100 to ensure
			// the items request sent to the getCommitData function is
			// accurate.
			items = branch.Compare.BehindBy % 100
			if items == 0 {
				items = 100
			}
		}
		c, err := ghs.getCommitData(ctx, client, repoName, items, cursor, branch.Name)
		if err != nil {
			ghs.logger.Sugar().Errorf("error making graphql query to get commit data", zap.Error(err))
			return 0, 0, 0, err
		}

		// TODO
		// let's make sure there's actually commits here stupid graphql nesting and types
		// if len(c.Edges) == 0 {
		if len(c.Nodes) == 0 {
			break
		}
		cursor = &c.PageInfo.EndCursor
		if page == pages {
			// e := c.GetEdges()
			node := c.GetNodes()
			oldest := node[len(node)-1].GetCommittedDate()
			age = int64(time.Since(oldest).Seconds())
		}
		for b := 0; b < len(c.Nodes); b++ {
			additions += c.Nodes[b].Additions
			deletions += c.Nodes[b].Deletions
		}

	}
	return additions, deletions, age, nil
}

func (ghs *githubScraper) getCommitData(
	ctx context.Context,
	client graphql.Client,
	repoName string,
	comCount int,
	cc *string,
	branchName string,
	// ) (*CommitNodeTargetCommitHistoryCommitHistoryConnection, error) {
) (*BranchHistoryTargetCommitHistoryCommitHistoryConnection, error) {
	data, err := getCommitData(ctx, client, repoName, ghs.cfg.GitHubOrg, 1, comCount, cc, branchName)
	if err != nil {
		return nil, err
	}

	// This checks to ensure that the query returned a BranchHistory Node. The
	// way the GraphQL query functions allows for a successful query to take
	// place, but have an empty set of branches. The only time this query would
	// return an empty BranchHistory Node is if the branch was deleted between
	// the time the list of branches was retrieved, and the query for the
	// commits on the branch.
	if len(data.Repository.Refs.Nodes) == 0 {
		return nil, errors.New("no branch history returned from the commit data request")
	}

	tar := data.Repository.Refs.Nodes[0].GetTarget()

	// i hate graphql types
	// if ct, ok := tar.(*CommitNodeTargetCommit); ok {
	if ct, ok := tar.(*BranchHistoryTargetCommit); ok {
		return &ct.History, nil
	}

	// TODO
	// this error message isn't accurate
	// i think this is the data requested for commits on a branch does not exist
	// well, it is actually a target that isn't a commit
	// this would be if the graphql query sent back something
	// target does not commit node target interface
	return nil, errors.New("target is not a commit")
}

func getNumPages(p float64, n float64) int {
	numPages := math.Ceil(n / p)

	return int(numPages)
}

func add[T ~int | ~float64](a, b T) T {
	return a + b
}

// Get the age/duration between two times in seconds.
func getAge(start time.Time, end time.Time) int64 {
	return int64(end.Sub(start).Seconds())
}
