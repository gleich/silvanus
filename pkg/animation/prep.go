package animation

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gleich/lumber/v2"
)

var (
	TempLocation  = filepath.Join(os.TempDir(), "silvanus")
	CloneLocation = filepath.Join(TempLocation, "repo")
)

// Prep the temp dir for cloning
func Prep(log lumber.Logger) error {
	fmt.Println()
	log.Info("Prepping temporary folder for making the animation")

	if _, err := os.Stat(TempLocation); !os.IsNotExist(err) {
		err := os.RemoveAll(TempLocation)
		if err != nil {
			return err
		}
		log.Info("Removed", TempLocation)
	}

	err := os.Mkdir(TempLocation, 0777)
	if err != nil {
		return err
	}
	log.Info("Created", TempLocation)

	err = os.Mkdir(CloneLocation, 0777)
	if err != nil {
		return err
	}
	log.Info("Created", CloneLocation)

	log.Success("Temporary folder for the animation setup")
	return nil
}
