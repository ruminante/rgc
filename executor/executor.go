package executor

import (
	"github.com/ruminante/conventional-commit/prompts"
	"github.com/spf13/cobra"
)

func Executor(cmd *cobra.Command) {

		if err := Add(cmd); err != nil {
			panic(err)
		}
			
		// Get Commit Type
		ctp := prompts.TypePrompt()
		i, _, err := ctp.Run(); if err != nil {
			panic(err)
		}

		if err := Msg(cmd, i); err != nil {
			panic(err)
		}
}