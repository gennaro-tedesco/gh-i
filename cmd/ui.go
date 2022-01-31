package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func parseInput(state string, title string, body string, user string, author string, labelsList []string) url.Values {
	queryString := fmt.Sprint("type:issue")
	if state != "" {
		queryString = queryString + fmt.Sprintf(" state:%s", state)
	}
	if title != "" {
		queryString = queryString + fmt.Sprintf(" %s in:title", title)
	}
	if body != "" {
		queryString = queryString + fmt.Sprintf(" %s in:body", body)
	}
	if user != "" {
		queryString = queryString + fmt.Sprintf(" user:%s", user)
	}
	queryString = queryString + fmt.Sprintf(" author:%s", author)
	for _, label := range labelsList {
		queryString = queryString + fmt.Sprintf(" label:%s", label)
	}
	query := url.Values{}
	query.Add("q", queryString)
	query.Add("per_page", "100")
	fmt.Println(queryString)
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
		Active:   fmt.Sprintf("\U0001F449 {{ .Title | %s | bold }}  {{ .State | faint }}", colour),
		Inactive: fmt.Sprintf("   {{ .Title | %s }}", colour),
		Selected: fmt.Sprintf(`{{ "✔" | green | bold }} {{ .Title | %s | bold }}`, colour),
		Details: `
	{{ "Title:" | faint }} 	{{ .Title }}
	{{ "Url address:" | faint }} 	{{ .URL }}
	`,
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
