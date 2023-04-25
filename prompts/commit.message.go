package prompts

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func MessagePrompt() *promptui.Prompt {
	validate := func(input string) error {
		len := len(input)
		if len > 255 {
			return errors.New("Message must have at most 255 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Write a commit message",
		Validate: validate,
	}

	return &prompt
}