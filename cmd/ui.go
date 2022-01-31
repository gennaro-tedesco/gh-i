package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func parseInput(filter string, state string, labelsList []string, sort string) url.Values {
	query := url.Values{}
	query.Add("filter", filter)
	query.Add("state", state)
	var labelsString string
	for _, label := range labelsList {
		labelsString = labelsString + fmt.Sprintf("%s,",label)
	}
	query.Add("sort", sort)
	query.Add("labels", labelsString)
	query.Add("per_page", "100")
	return query
}

func getTemplate(colour string) *promptui.SelectTemplates {
	funcMap := promptui.FuncMap

	funcMap["truncate"] = func(input string) string {
		length := 80
		if len(input) <= length {
			return input
		}
		return input[:length-3] + "..."
	}

	return &promptui.SelectTemplates{
		Active:   fmt.Sprintf("\U0001F449 {{ .Title | %s | bold }}", colour),
		Inactive: fmt.Sprintf("   {{ .Title | %s }}", colour),
		Selected: fmt.Sprintf(`{{ "✔" | green | bold }} {{ .Title | %s | bold }}`, colour),
		Details: `
	{{ "Title:" | faint }} 	{{ .Title }}
	{{ "Url address:" | faint }} 	{{ .URL }}
	{{ "UpdatedAt" | faint }} 	{{ .UpdatedAt }}`,
	}

}

func getSelectionPrompt(issues []issueInfo, colour string) *promptui.Select {
	return &promptui.Select{
		Stdout:    os.Stderr,
		Stdin:     os.Stdin,
		Label:     "repository list",
		Items:     issues,
		Templates: getTemplate(colour),
		Size:      20,
		Searcher: func(input string, idx int) bool {
			repo := issues[idx]
			title := strings.ToLower(repo.Title)

			return strings.Contains(title, input)
		},
	}
}

