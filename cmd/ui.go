package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func parseInput(state string, title string, body string, user string, me bool, labelsList []string) url.Values {
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
	if me {
		queryString = queryString + fmt.Sprintf(" author:@me")
	}
	for _, label := range labelsList {
		queryString = queryString + fmt.Sprintf(" label:%s", label)
	}
	query := url.Values{}
	query.Add("q", queryString)
	query.Add("sort", "updated")
	query.Add("per_page", "100")
	return query
}

func getTemplate(colour string) *promptui.SelectTemplates {
	funcMap := promptui.FuncMap
	funcMap["parseLabels"] = func(Labels []interface{}) string {
		if len(Labels) == 0 {
			return ""
		}
		var concatLabels string
		for _, label := range Labels {
			concatLabels = concatLabels + fmt.Sprintf("%s, ", label.(map[string]interface{})["name"].(string))
		}
		return concatLabels[:len(concatLabels)-2]
	}

	return &promptui.SelectTemplates{
		Active:   fmt.Sprintf("\U0001F449 {{ .Title | %s | bold }}  {{ .State | faint }}", colour),
		Inactive: fmt.Sprintf("   {{ .Title | %s }}", colour),
		Selected: fmt.Sprintf(`{{ "✔" | green | bold }} {{ .Title | %s | bold }}`, colour),
		Details: `
	{{ "Title:" | faint }} 	{{ .Title }}
	{{ "Url address:" | faint }} 	{{ .URL }}
	{{ "Labels:" | faint }} 	{{ .Labels | parseLabels }}
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
