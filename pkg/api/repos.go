package api

import (
	"context"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/silvanus/pkg/ask"
	"github.com/shurcooL/githubv4"
)

// Outline for a repo from GitHub
type Repo struct {
	Description   string
	NameWithOwner string
	URL           string
}

// Get repos that belong to the user
func Repos(log lumber.Logger, client *githubv4.Client, opts ask.Opts) ([]Repo, error) {
	log.Info("Getting repos you belong to")

	repos := []Repo{}
	total := 0
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

		new := len(data.Viewer.Repositories.Nodes)
		total += new
		log.Success("Added", new, "repos. Total is now", total)

		if !data.Viewer.Repositories.PageInfo.HasNextPage {
			log.Info("Reached final page of data")
			break
		}
	}

	log.Success("Got data for", total, "repos")
	return repos, nil
}
