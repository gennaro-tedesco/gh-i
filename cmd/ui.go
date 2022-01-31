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
		labelsString = labelsString + fmt.Sprintf(label,",")
	}
	query.Add("sort", sort)
	query.Add("per_page", "100")
	return query
}

func getTemplate(colour string) *promptui.SelectTemplates {
	funcMap := promptui.FuncMap
	funcMap["parseStars"] = func(starCount float64) string {
		if starCount >= 1000 {
			return fmt.Sprintf("%.1f k", starCount/1000)
		}
		return fmt.Sprint(starCount)
	}

	funcMap["truncate"] = func(input string) string {
		length := 80
		if len(input) <= length {
			return input
		}
		return input[:length-3] + "..."
	}

	return &promptui.SelectTemplates{
		Active:   fmt.Sprintf("\U0001F449 {{ .Name | %s | bold }}", colour),
		Inactive: fmt.Sprintf("   {{ .Name | %s }}", colour),
		Selected: fmt.Sprintf(`{{ "✔" | green | bold }} {{ .Name | %s | bold }}`, colour),
		Details: `
	{{ "Name:" | faint }} 	 {{ .Name }}
	{{ "Description:" | faint }} 	 {{ .Description | truncate }}
	{{ "Url address:" | faint }} 	 {{ .URL }}
	{{ "⭐" | faint }}	{{ .Stars | parseStars }}`,
	}

}

func getSelectionPrompt(repos []repoInfo, colour string) *promptui.Select {
	return &promptui.Select{
		Stdout:    os.Stderr,
		Stdin:     os.Stdin,
		Label:     "repository list",
		Items:     repos,
		Templates: getTemplate(colour),
		Size:      20,
		Searcher: func(input string, idx int) bool {
			repo := repos[idx]
			title := strings.ToLower(repo.Name)

			return strings.Contains(title, input)
		},
	}
}

