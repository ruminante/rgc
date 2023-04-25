package prompts

import (
	"strings"

	"github.com/manifoldco/promptui"
)

type commit_types struct {
	Name     string
	Description  string
	Value string
}

var Types = []commit_types{
	{Name: "Feat", Description: "A new feature", Value: "feat: :sparkles: "},
	{Name: "Fix", Description: "A bug fix", Value: "fix: :bug: "},
	{Name: "Docs", Description: "Changes to Documentation", Value: "docs: :memo: "},
	{Name: "Style", Description: "Changes that do not affect the meaning of the code (Ex: Indentation, semi-colons, etc...)", Value: "style: :lipstick: "},
	{Name: "Refactor", Description: "A Change that doesn't fix a bug nor adds a feature", Value: "refactor: :art: "},
	{Name: "Performance", Description: "A code change that improves performance", Value: "perf: :zap: "},
	{Name: "Test", Description: "Adds a missing test or correct an existing test", Value: "test: :test_tube: "},
	{Name: "Build", Description: "Changes that affect the build or external dependencies", Value: "build: :construction: "},
	{Name: "CI", Description: "Changes to our CI configuration (Ex: Github Workflows, Husky, etc...)", Value: "ci: :construction_worker: "},
	{Name: "Chore", Description: "Other changes that don't modify src or test files", Value: "chore: :wrench: "},
	{Name: "Revert", Description: "Reverts a previous commit", Value: "revert: :rewind: "},
}

func TypePrompt() promptui.Select {	

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "> {{ .Name | green }} ",
		Inactive: "  {{ .Name | red }} ",
		Selected: "> {{ .Name | red | cyan }}",
		Details: `description: {{ .Description | faint }}`,
	}

	searcher := func(input string, index int) bool {
		cm_type := Types[index]
		name := strings.Replace(strings.ToLower(cm_type.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select a commit type",
		Items:     Types,
		Templates: templates,
		Size:      5,
		Searcher:  searcher,
	}

	return prompt
}