package ask

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

// Options
type Opts struct {
	Private bool
	Max     int
}

// Ask user questions surrounding the avaliable options
func Options() (Opts, error) {
	questions := []*survey.Question{
		{
			Name:   "private",
			Prompt: &survey.Confirm{Message: "Should private repos should be included?"},
		},
		{
			Name: "max",
			Prompt: &survey.Input{
				Message: "What is the max number of repos included?",
				Default: "1000",
			},
			Validate: func(ans interface{}) error {
				_, err := strconv.ParseInt(ans.(string), 10, 64)
				if err != nil {
					return errors.New(fmt.Sprint(ans, " isn't a valid number"))
				}
				return nil
			},
		},
	}
	var answers Opts
	err := survey.Ask(questions, &answers)
	if err != nil {
		return Opts{}, err
	}
	return answers, nil
}
