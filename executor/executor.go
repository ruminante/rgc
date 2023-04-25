package executor

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ruminante/conventional-commit/prompts"
	"github.com/spf13/cobra"
)

func Executor(cmd *cobra.Command) {

	if cmd.Flags().Lookup("logs").Changed {
		test := exec.Command("git", "log", "--pretty=format:%C(blue)%h%C(red)%d %C(white)%s - %C(cyan)%cn, %C(green)%cr")
		test.Stdout = os.Stdout
		if err := test.Run(); err != nil {
			fmt.Printf("An error occured during the commit log, reason: %s \n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if cmd.Flags().Lookup("undo").Changed {
		if err := exec.Command("git", "reset", "HEAD~1", "--soft").Run(); err != nil {
			fmt.Printf("An error occured during the commit revert, reason: %s \n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

		if err := Add(cmd); err != nil {
			fmt.Printf("An error occured during the Add files prompt, reason: %s \n", err)
			os.Exit(1)
		}
			
		// Get Commit Type
		ctp := prompts.TypePrompt()
		i, _, err := ctp.Run(); if err != nil {
			fmt.Printf("An error occured during the commit type prompt, reason: %s \n", err)
			os.Exit(1)
		}

		if err := Msg(cmd, i); err != nil {
			fmt.Printf("An error occured during the commit message prompt, reason: %s \n", err)
			os.Exit(1)
		}
}