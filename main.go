package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/silvanus/pkg/animation"
	"github.com/gleich/silvanus/pkg/api"
	"github.com/gleich/silvanus/pkg/ask"
)

func main() {
	log := lumber.NewCustomLogger()
	log.Timezone = time.Local

	token, err := ask.Token()
	if err != nil {
		log.Fatal(err, "Failed to ask for token")
	}

	opts, err := ask.Options()
	if err != nil {
		log.Fatal(err, "Failed to ask options")
	}

	client := api.NewClient(token)

	repos, err := api.Repos(log, client, opts)
	if err != nil {
		log.Fatal(err, "Failed to get data for repos")
	}

	err = ask.ConfirmGen(log, len(repos), animation.CloneLocation)
	if err != nil {
		log.Fatal(err, "Failed to confirm with user about generation of animation")
	}

	err = animation.Prep(log)
	if err != nil {
		log.Fatal(err, "Failed to prep temp dir for creating the animation")
	}

	for i := 0; i < len(repos); i++ {
		repo := repos[i]
		fmt.Println()
		log.Info("Generating log for", repo.NameWithOwner, fmt.Sprint(i+1, "/", len(repos)))

		gitPath, err := exec.LookPath("git")
		if err != nil {
			log.Fatal(
				err,
				"Failed to locate the git executable. Please make sure it's in your PATH",
			)
		}

		err = animation.Clone(log, repo, gitPath)
		if err != nil {
			log.Fatal(err, "Failed to clone", repo.NameWithOwner)
		}
		break
	}
}
