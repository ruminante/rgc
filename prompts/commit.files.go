package prompts

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

type commit_files struct {
	Path     string
	IsSelected bool
}

func getPaths(bts []byte) []string {
	rmv_nl := strings.Replace(string(bts), "\n", "", -1)
	rmv_intr := strings.Replace(string(rmv_nl), "??", "", -1)
	rmv_m := strings.Replace(string(rmv_intr), "M", "", -1)
	rmv_d := strings.Replace(string(rmv_m), "D", "", -1)
	paths := strings.Split(string(rmv_d), " ")
	paths[0] = "."
	for i, path := range paths {
		if path == "" {
			paths = append(paths[:i], paths[i+1:]...)
		}
	}
	paths = append(paths, "Done")
	return paths
}

func GetList() ([]*commit_files, error) {
	cmd := exec.Command("git", "status", "-s", "-u")
	out, err := cmd.Output(); if err != nil {
		return nil, err
	}
	paths := getPaths(out)

	files := make([]*commit_files, (len(paths)))

	for i, path := range paths {
		files[i] = &commit_files{Path: path}
	}
	return files, nil
}

func SelectItems(selectedPos int, allItems []*commit_files) ([]*commit_files, error) {

	templates := &promptui.SelectTemplates{
			Label: `{{if .IsSelected}}
									✔
							{{end}} {{ .Path }} - label`,
			Active:   "{{if .IsSelected}}✔ {{end}}{{ .Path | green }}",
			Inactive: "{{if .IsSelected}}✔ {{end}}{{ .Path | red }}",
	}

	searcher := func(input string, index int) bool {
		cm_files := allItems[index]
		name := strings.Replace(strings.ToLower(cm_files.Path), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}
	
	prompt := promptui.Select{
			Label:     "Item",
			Items:     allItems,
			Templates: templates,
			Size:      len(allItems),
			CursorPos:    selectedPos,
			HideSelected: true,
			Searcher: searcher,
	}
	
	selectionIdx, _, err := prompt.Run()
			if err != nil {
			return nil, fmt.Errorf("prompt failed: %w", err)
	}
	
	chosenItem := allItems[selectionIdx]
	
	if chosenItem.Path != "Done" && chosenItem.Path != "." {
			chosenItem.IsSelected = !chosenItem.IsSelected
			return SelectItems(selectionIdx, allItems)
	}

	if chosenItem.Path == "." {
		chosenItem.IsSelected = true
	}
	
	var selectedItems []*commit_files
	for _, i := range allItems {
			if i.IsSelected {
					selectedItems = append(selectedItems, i)
			}
	}

	return selectedItems, nil
}