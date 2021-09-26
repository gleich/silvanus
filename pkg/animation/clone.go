package animation

import (
	"os"
	"os/exec"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/silvanus/pkg/api"
)

func Clone(log lumber.Logger, repo api.Repo, gitPath string) error {
	log.Info("Cloning")
	cmd := &exec.Cmd{
		Path: gitPath,
		Args: []string{
			gitPath,
			"clone",
			repo.URL + ".git",
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Dir:    CloneLocation,
	}
	err := cmd.Run()
	if err != nil {
		return err
	}
	log.Success("Cloned")
	return nil
}
