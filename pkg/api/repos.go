package api

import (
	"context"
	"fmt"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/silvanus/pkg/ask"
	"github.com/shurcooL/githubv4"
)

// Outline for a repo from GitHub
type Repo struct {
	Description   string
	Name          string
	NameWithOwner string
	URL           string
}

// Get repos that belong to the user
func Repos(log lumber.Logger, client *githubv4.Client, opts ask.Opts) ([]Repo, error) {
	fmt.Println()
	log.Info("Getting repos you belong to")

	repos := []Repo{}
	variables := map[string]interface{}{
		"cursor":  (*githubv4.String)(nil),
		"privacy": (*githubv4.RepositoryPrivacy)(nil),
	}
	if !opts.Private {
		variables["privacy"] = githubv4.RepoAccessAuditEntryVisibilityPublic
	}

	for {
		var data struct {
			Viewer struct {
				Repositories struct {
					Nodes    []Repo
					PageInfo struct {
						HasNextPage bool
						EndCursor   githubv4.String
					}
				} `graphql:"repositories(first: 100, privacy: $privacy, after: $cursor)"`
			}
		}
		err := client.Query(context.Background(), &data, variables)
		if err != nil {
			return []Repo{}, err
		}
		variables["cursor"] = githubv4.NewString(data.Viewer.Repositories.PageInfo.EndCursor)

		repos = append(repos, data.Viewer.Repositories.Nodes...)
		log.Info("Added", len(data.Viewer.Repositories.Nodes), "repos. Total is now", len(repos))

		if !data.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
	}

	log.Success("Got data for", len(repos), "repos")
	return repos, nil
}
