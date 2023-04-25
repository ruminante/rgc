package executor

import (
	"fmt"
	"os"

	"github.com/ruminante/conventional-commit/prompts"
	"github.com/spf13/cobra"
)

func Executor(cmd *cobra.Command) {

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