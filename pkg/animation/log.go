package animation

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/silvanus/pkg/api"
)

func GenLog(log lumber.Logger, repo api.Repo, i int, gourcePath string) error {
	log.Info("Generating log")
	cmd := &exec.Cmd{
		Path: gourcePath,
		Args: []string{
			gourcePath,
			"--output-custom-log",
			fmt.Sprintf("log%v.txt", i),
			filepath.Join("repo", repo.Name),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Dir:    TempLocation,
	}
	err := cmd.Run()
	if err != nil {
		return err
	}
	log.Success("Generated log")
	return nil
}
