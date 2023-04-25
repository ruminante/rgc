package executor

import (
	"os/exec"

	"github.com/ruminante/conventional-commit/prompts"
	"github.com/spf13/cobra"
)

func Msg(cmd *cobra.Command, i int) error {
	if cmd.Flags().Lookup("message").Changed {
		m, _ := cmd.Flags().GetString("message")
		err := exec.Command("git", "commit", "-m", prompts.Types[i].Value+m).Run()
		if err != nil {return err}
		return nil
	}
	msg, err := prompts.MessagePrompt().Run(); if err != nil {
		return err
	}
	if exec.Command("git", "commit", "-m", prompts.Types[i].Value+msg).Run(); err != nil {return err}
	return nil
}