package ask

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gleich/lumber/v2"
	"github.com/gleich/silvanus/pkg/animation"
)

// Confirm with the user that they want to generate a gource animation using the number of repos provided
func ConfirmGen(log lumber.Logger, repos int) error {
	fmt.Println()
	var answer bool
	err := survey.AskOne(
		&survey.Confirm{
			Message: fmt.Sprintf(
				`Are you sure that you want to generate a gource animation for %v repos?

The following will happen for each repo:
	1. Cloned in %[2]v.
	2. Gource log generated for it
	3. Removed from %[2]v

This process could take a few minutes to a few hours depending on how many repos are being cloned and how big they are.`,
				repos,
				animation.CloneLocation,
			),
		},
		&answer,
	)
	if err != nil {
		return err
	}

	if !answer {
		log.Info("Generation cancelled by user. Have a good day :)")
		os.Exit(0)
	}

	return nil
}
