package prompts

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
)

type commit_files struct {
	Path     string
	IsSelected bool
}

func getPaths(bts []byte) []string {
	re := regexp.MustCompile("[\n??MD]")
	bts = re.ReplaceAll(bts, []byte(""))
	paths := make([]string, 0, bytes.Count(bts, []byte(" "))+2)
	fields := bytes.Fields(bts)

	for _, field := range fields {
		paths = append(paths, string(field))
	}

	paths = append(paths, "all files", "finish")
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

	if chosenItem.Path == "all files" {
		for i := range allItems {
			if allItems[i].Path != "finish" {
				allItems[i].IsSelected = true
			}
		}
	}
	
	if chosenItem.Path != "finish" {
			chosenItem.IsSelected = !chosenItem.IsSelected
			return SelectItems(selectionIdx, allItems)
	}

	
	var selectedItems []*commit_files
	for _, i := range allItems {
			if i.IsSelected {
					selectedItems = append(selectedItems, i)
			}
	}

	return selectedItems, nil
}