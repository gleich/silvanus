package ask

import "github.com/AlecAivazis/survey/v2"

// Ask the user for their GitHub PAT
func Token() (string, error) {
	token := ""
	err := survey.AskOne(&survey.Password{Message: "GitHub token"}, &token)
	if err != nil {
		return "", err
	}
	return token, nil
}
