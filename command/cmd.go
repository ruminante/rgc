package command

import (
	"fmt"
	"os"

	"github.com/ruminante/conventional-commit/executor"
	"github.com/spf13/cobra"
)

var undo bool
var log bool
var add bool
var m string

var rootCmd = &cobra.Command{
  Use:   "rgc [flags]",
  Short: "rgc is a conventional commit tool",
  Long: `rgc (Ruminante Conventional Commit) is a conventional commit tool to help you write better commit messages and follow the conventional commit standard.`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
		executor.Executor(cmd)
  },
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&add, "all", "a", false, "Select all files for git add")
	rootCmd.PersistentFlags().StringVarP(&m, "message", "m", "", "Commit message for the commit")
  rootCmd.PersistentFlags().BoolVarP(&log, "logs", "l", false, "Logs all commits in the repository")
	rootCmd.PersistentFlags().BoolVar(&undo, "undo", false, "Select all files for git add")
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}