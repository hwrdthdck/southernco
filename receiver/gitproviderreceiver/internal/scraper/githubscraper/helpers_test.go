// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package githubscraper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/google/go-github/v57/github"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/receiver/receivertest"
)

type responses struct {
	repoResponse       repoResponse
	prResponse         prResponse
	branchResponse     branchResponse
	checkLoginResponse loginResponse
	contribResponse    contribResponse
	scrape             bool
}

type repoResponse struct {
	repos        []getRepoDataBySearchSearchSearchResultItemConnection
	responseCode int
	page         int
}

type prResponse struct {
	prs          []getPullRequestDataRepositoryPullRequestsPullRequestConnection
	responseCode int
	page         int
}

type branchResponse struct {
	branches     []getBranchDataRepositoryRefsRefConnection
	responseCode int
	page         int
}

type loginResponse struct {
	checkLogin   checkLoginResponse
	responseCode int
}

type contribResponse struct {
	contribs     [][]*github.Contributor
	responseCode int
	page         int
}

func MockServer(responses *responses) *http.ServeMux {
	var mux http.ServeMux
	restEndpoint := "/api-v3/repos/o/r/contributors"
	graphEndpoint := "/"
	if responses.scrape {
		graphEndpoint = "/api/graphql"
		restEndpoint = "/api/v3/repos/liatrio/repo1/contributors"
	}
	mux.HandleFunc(graphEndpoint, func(w http.ResponseWriter, r *http.Request) {
		var reqBody graphql.Request
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			return
		}
		switch {
		// These OpNames need to be name of the GraphQL query as defined in genqlient.graphql
		case reqBody.OpName == "checkLogin":
			loginResp := &responses.checkLoginResponse
			w.WriteHeader(loginResp.responseCode)
			if loginResp.responseCode == http.StatusOK {
				login := loginResp.checkLogin
				graphqlResponse := graphql.Response{Data: &login}
				if err := json.NewEncoder(w).Encode(graphqlResponse); err != nil {
					return
				}
			}
		case reqBody.OpName == "getRepoDataBySearch":
			repoResp := &responses.repoResponse
			w.WriteHeader(repoResp.responseCode)
			if repoResp.responseCode == http.StatusOK {
				repos := getRepoDataBySearchResponse{
					Search: repoResp.repos[repoResp.page],
				}
				graphqlResponse := graphql.Response{Data: &repos}
				if err := json.NewEncoder(w).Encode(graphqlResponse); err != nil {
					return
				}
				repoResp.page++
			}
		case reqBody.OpName == "getBranchData":
			branchResp := &responses.branchResponse
			w.WriteHeader(branchResp.responseCode)
			if branchResp.responseCode == http.StatusOK {
				branches := getBranchDataResponse{
					Repository: getBranchDataRepository{
						Refs: branchResp.branches[branchResp.page],
					},
				}
				graphqlResponse := graphql.Response{Data: &branches}
				if err := json.NewEncoder(w).Encode(graphqlResponse); err != nil {
					return
				}
				branchResp.page++
			}
		case reqBody.OpName == "getPullRequestData":
			prResp := &responses.prResponse
			w.WriteHeader(prResp.responseCode)
			if prResp.responseCode == http.StatusOK {
				repos := getPullRequestDataResponse{
					Repository: getPullRequestDataRepository{
						PullRequests: prResp.prs[prResp.page],
					},
				}
				graphqlResponse := graphql.Response{Data: &repos}
				if err := json.NewEncoder(w).Encode(graphqlResponse); err != nil {
					return
				}
				prResp.page++
			}
		}
	})
	mux.HandleFunc(restEndpoint, func(w http.ResponseWriter, r *http.Request) {
		contribResp := &responses.contribResponse
		if contribResp.responseCode == http.StatusOK {
			contribs, err := json.Marshal(contribResp.contribs[contribResp.page])
			if err != nil {
				fmt.Printf("error marshaling response: %v", err)
			}
			link := fmt.Sprintf(
				"<https://api.github.com/repositories/placeholder/contributors?per_page=100&page=%d>; rel=\"next\"",
				len(contribResp.contribs)-contribResp.page-1,
			)
			w.Header().Set("Link", link)
			// Attempt to write data to the response writer.
			_, err = w.Write(contribs)
			if err != nil {
				fmt.Printf("error writing response: %v", err)
			}
			contribResp.page++
		}
	})
	return &mux
}

func TestGenDefaultSearchQueryOrg(t *testing.T) {
	st := "org"
	org := "empire"

	expected := "org:empire archived:false"

	actual := genDefaultSearchQuery(st, org)

	assert.Equal(t, expected, actual)
}

func TestGenDefaultSearchQueryUser(t *testing.T) {
	st := "user"
	org := "vader"

	expected := "user:vader archived:false"

	actual := genDefaultSearchQuery(st, org)

	assert.Equal(t, expected, actual)
}

func TestGetAge(t *testing.T) {
	testCases := []struct {
		desc     string
		hrsAdd   time.Duration
		minsAdd  time.Duration
		expected float64
	}{
		{
			desc:     "TestHalfHourDiff",
			hrsAdd:   0 * time.Hour,
			minsAdd:  time.Duration(30) * time.Minute,
			expected: 60 * 30,
		},
		{
			desc:     "TestHourDiff",
			hrsAdd:   1 * time.Hour,
			minsAdd:  0 * time.Minute,
			expected: 60 * 60,
		},
		{
			desc:     "TestDayDiff",
			hrsAdd:   24 * time.Hour,
			minsAdd:  0 * time.Minute,
			expected: 60 * 60 * 24,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			min := time.Now()
			max := min.Add(tc.hrsAdd).Add(tc.minsAdd)

			actual := getAge(min, max)

			assert.Equal(t, int64(tc.expected), actual)
		})
	}
}

func TestCheckOwnerExists(t *testing.T) {
	testCases := []struct {
		desc              string
		login             string
		expectedError     bool
		expectedOwnerType string
		server            *http.ServeMux
	}{
		{
			desc:  "TestOrgOwnerExists",
			login: "liatrio",
			server: MockServer(&responses{
				checkLoginResponse: loginResponse{
					checkLogin: checkLoginResponse{
						Organization: checkLoginOrganization{
							Login: "liatrio",
						},
					},
					responseCode: http.StatusOK,
				},
			}),
			expectedOwnerType: "org",
		},
		{
			desc:  "TestUserOwnerExists",
			login: "liatrio",
			server: MockServer(&responses{
				checkLoginResponse: loginResponse{
					checkLogin: checkLoginResponse{
						User: checkLoginUser{
							Login: "liatrio",
						},
					},
					responseCode: http.StatusOK,
				},
			}),
			expectedOwnerType: "user",
		},
		{
			desc:  "TestLoginError",
			login: "liatrio",
			server: MockServer(&responses{
				checkLoginResponse: loginResponse{
					checkLogin: checkLoginResponse{
						Organization: checkLoginOrganization{
							Login: "liatrio",
						},
					},
					responseCode: http.StatusNotFound,
				},
			}),
			expectedOwnerType: "",
			expectedError:     true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			factory := Factory{}
			defaultConfig := factory.CreateDefaultConfig()
			settings := receivertest.NewNopCreateSettings()
			ghs := newGitHubScraper(context.Background(), settings, defaultConfig.(*Config))
			server := httptest.NewServer(tc.server)
			defer server.Close()

			client := graphql.NewClient(server.URL, ghs.client)
			loginType, err := ghs.login(context.Background(), client, tc.login)

			assert.Equal(t, tc.expectedOwnerType, loginType)
			if !tc.expectedError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetPullRequests(t *testing.T) {
	testCases := []struct {
		desc            string
		server          *http.ServeMux
		expectedErr     error
		expectedPrCount int
	}{
		{
			desc: "TestSinglePageResponse",
			server: MockServer(&responses{
				scrape: false,
				prResponse: prResponse{
					prs: []getPullRequestDataRepositoryPullRequestsPullRequestConnection{
						{
							PageInfo: getPullRequestDataRepositoryPullRequestsPullRequestConnectionPageInfo{
								HasNextPage: false,
							},
							Nodes: []PullRequestNode{
								{
									Merged: false,
								},
								{
									Merged: false,
								},
								{
									Merged: false,
								},
							},
						},
					},
					responseCode: http.StatusOK,
				},
			}),
			expectedErr:     nil,
			expectedPrCount: 3, // 3 PRs per page, 1 pages
		},
		{
			desc: "TestMultiPageResponse",
			server: MockServer(&responses{
				scrape: false,
				prResponse: prResponse{
					prs: []getPullRequestDataRepositoryPullRequestsPullRequestConnection{
						{
							PageInfo: getPullRequestDataRepositoryPullRequestsPullRequestConnectionPageInfo{
								HasNextPage: true,
							},
							Nodes: []PullRequestNode{
								{
									Merged: false,
								},
								{
									Merged: false,
								},
								{
									Merged: false,
								},
							},
						},
						{
							PageInfo: getPullRequestDataRepositoryPullRequestsPullRequestConnectionPageInfo{
								HasNextPage: false,
							},
							Nodes: []PullRequestNode{
								{
									Merged: false,
								},
								{
									Merged: false,
								},
								{
									Merged: false,
								},
							},
						},
					},
					responseCode: http.StatusOK,
				},
			}),
			expectedErr:     nil,
			expectedPrCount: 6, // 3 PRs per page, 2 pages
		},
		{
			desc: "Test404Response",
			server: MockServer(&responses{
				scrape: false,
				prResponse: prResponse{
					responseCode: http.StatusNotFound,
				},
			}),
			expectedErr:     errors.New("returned error 404 Not Found: "),
			expectedPrCount: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			factory := Factory{}
			defaultConfig := factory.CreateDefaultConfig()
			settings := receivertest.NewNopCreateSettings()
			ghs := newGitHubScraper(context.Background(), settings, defaultConfig.(*Config))
			server := httptest.NewServer(tc.server)
			defer server.Close()
			client := graphql.NewClient(server.URL, ghs.client)

			prs, err := ghs.getPullRequests(context.Background(), client, "repo name")

			assert.Equal(t, tc.expectedPrCount, len(prs))
			if tc.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr.Error())
			}
		})
	}
}

func TestGetRepos(t *testing.T) {
	testCases := []struct {
		desc        string
		server      *http.ServeMux
		expectedErr error
		expected    int
	}{
		{
			desc: "TestSinglePageResponse",
			server: MockServer(&responses{
				repoResponse: repoResponse{
					repos: []getRepoDataBySearchSearchSearchResultItemConnection{
						{
							RepositoryCount: 1,
							Nodes: []SearchNode{
								&SearchNodeRepository{
									Name: "repo1",
								},
							},
							PageInfo: getRepoDataBySearchSearchSearchResultItemConnectionPageInfo{
								HasNextPage: false,
							},
						},
					},
					responseCode: http.StatusOK,
				},
			}),
			expectedErr: nil,
			expected:    1,
		},
		{
			desc: "TestMultiPageResponse",
			server: MockServer(&responses{
				repoResponse: repoResponse{
					repos: []getRepoDataBySearchSearchSearchResultItemConnection{
						{
							RepositoryCount: 4,
							Nodes: []SearchNode{
								&SearchNodeRepository{
									Name: "repo1",
								},
								&SearchNodeRepository{
									Name: "repo2",
								},
							},
							PageInfo: getRepoDataBySearchSearchSearchResultItemConnectionPageInfo{
								HasNextPage: true,
							},
						},
						{
							RepositoryCount: 4,
							Nodes: []SearchNode{
								&SearchNodeRepository{
									Name: "repo3",
								},
								&SearchNodeRepository{
									Name: "repo4",
								},
							},
							PageInfo: getRepoDataBySearchSearchSearchResultItemConnectionPageInfo{
								HasNextPage: false,
							},
						},
					},
					responseCode: http.StatusOK,
				},
			}),
			expectedErr: nil,
			expected:    4,
		},
		{
			desc: "Test404Response",
			server: MockServer(&responses{
				repoResponse: repoResponse{
					responseCode: http.StatusNotFound,
				},
			}),
			expectedErr: errors.New("returned error 404 Not Found: "),
			expected:    0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			factory := Factory{}
			defaultConfig := factory.CreateDefaultConfig()
			settings := receivertest.NewNopCreateSettings()
			ghs := newGitHubScraper(context.Background(), settings, defaultConfig.(*Config))
			server := httptest.NewServer(tc.server)
			defer server.Close()
			client := graphql.NewClient(server.URL, ghs.client)

			_, count, err := ghs.getRepos(context.Background(), client, "fake query")

			assert.Equal(t, tc.expected, count)
			if tc.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr.Error())
			}
		})
	}
}

func TestGetBranches(t *testing.T) {
	testCases := []struct {
		desc        string
		server      *http.ServeMux
		expectedErr error
		expected    int
	}{
		{
			desc: "TestSinglePageResponse",
			server: MockServer(&responses{
				branchResponse: branchResponse{
					branches: []getBranchDataRepositoryRefsRefConnection{
						{
							TotalCount: 1,
							Nodes: []BranchNode{
								{
									Name: "main",
								},
							},
							PageInfo: getBranchDataRepositoryRefsRefConnectionPageInfo{
								HasNextPage: false,
							},
						},
					},
					responseCode: http.StatusOK,
				},
			}),
			expectedErr: nil,
			expected:    1,
		},
		{
			desc: "TestMultiPageResponse",
			server: MockServer(&responses{
				branchResponse: branchResponse{
					branches: []getBranchDataRepositoryRefsRefConnection{
						{
							TotalCount: 4,
							Nodes: []BranchNode{
								{
									Name: "main",
								},
								{
									Name: "vader",
								},
							},
							PageInfo: getBranchDataRepositoryRefsRefConnectionPageInfo{
								HasNextPage: true,
							},
						},
						{
							TotalCount: 4,
							Nodes: []BranchNode{
								{
									Name: "skywalker",
								},
								{
									Name: "rebelalliance",
								},
							},
							PageInfo: getBranchDataRepositoryRefsRefConnectionPageInfo{
								HasNextPage: false,
							},
						},
					},
					responseCode: http.StatusOK,
				},
			}),
			expectedErr: nil,
			expected:    4,
		},
		{
			desc: "Test404Response",
			server: MockServer(&responses{
				branchResponse: branchResponse{
					responseCode: http.StatusNotFound,
				},
			}),
			expectedErr: errors.New("returned error 404 Not Found: "),
			expected:    0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			factory := Factory{}
			defaultConfig := factory.CreateDefaultConfig()
			settings := receivertest.NewNopCreateSettings()
			ghs := newGitHubScraper(context.Background(), settings, defaultConfig.(*Config))
			server := httptest.NewServer(tc.server)
			defer server.Close()
			client := graphql.NewClient(server.URL, ghs.client)

			count, err := ghs.getBranches(context.Background(), client, "deathstarrepo", "main")

			assert.Equal(t, tc.expected, count)
			if tc.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr.Error())
			}
		})
	}
}

func TestGetContributors(t *testing.T) {
	testCases := []struct {
		desc          string
		server        *http.ServeMux
		repo          string
		org           string
		expectedErr   error
		expectedCount int
	}{
		{
			desc: "TestListContributorsResponse",
			server: MockServer(&responses{
				contribResponse: contribResponse{
					contribs: [][]*github.Contributor{{
						{
							ID: github.Int64(1),
						},
						{
							ID: github.Int64(2),
						},
					}},
					responseCode: http.StatusOK,
				},
			}),
			repo:          "r",
			org:           "o",
			expectedErr:   nil,
			expectedCount: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			factory := Factory{}
			defaultConfig := factory.CreateDefaultConfig()
			settings := receivertest.NewNopCreateSettings()
			ghs := newGitHubScraper(context.Background(), settings, defaultConfig.(*Config))
			ghs.cfg.GitHubOrg = tc.org

			server := httptest.NewServer(tc.server)

			client := github.NewClient(nil)
			url, _ := url.Parse(server.URL + "/api-v3" + "/")
			client.BaseURL = url
			client.UploadURL = url

			contribs, err := ghs.getContributorCount(context.Background(), client, tc.repo)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCount, contribs)
		})
	}
}
