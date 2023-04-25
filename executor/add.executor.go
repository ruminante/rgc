package executor

import (
	"fmt"
	"os/exec"

	"github.com/ruminante/conventional-commit/prompts"
	"github.com/spf13/cobra"
)

func Add(cmd *cobra.Command) error {

	if cmd.Flags().Lookup("all").Changed {
		cmd := exec.Command("git", "add", ".")
		_, err := cmd.Output(); if err != nil {
			return fmt.Errorf("Error adding files, check if you are in a git repository")
		}
		return nil
	}

	file_list, err := prompts.GetList(); if err != nil {
		return fmt.Errorf("Error adding files, check if you are in a git repository")
	}

	selected_files, err := prompts.SelectItems(0, file_list); if err != nil {
		return fmt.Errorf("%s", err)
	}

	for _, fs := range selected_files {
		cmd := exec.Command("git", "add", fs.Path)
		_, err := cmd.Output(); if err != nil {return fmt.Errorf("Error adding files, check if you are in a git repository")}
	}
	return nil
}