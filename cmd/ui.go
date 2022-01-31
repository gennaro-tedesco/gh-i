package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/manifoldco/promptui"
	"golang.org/x/term"
)

func colourMap() map[string]text.Color {
	colourMap := map[string]text.Color{
		"black":   text.FgBlack,
		"cyan":    text.FgCyan,
		"green":   text.FgGreen,
		"yellow":  text.FgYellow,
		"blue":    text.FgBlue,
		"magenta": text.FgMagenta,
		"red":     text.FgRed,
		"white":   text.FgWhite,
	}
	return colourMap
}

func explainInput(state string, title string, body string, user string, me bool, labelsList []string, colour string) {
	var queryString string
	if state == "" {
		queryString = queryString + fmt.Sprintf(" state:any +")
	} else {
		queryString = queryString + fmt.Sprintf(" state:%s +", state)
	}
	if me {
		queryString = queryString + fmt.Sprintf(" author:yourself +")
	} else {
		queryString = queryString + fmt.Sprintf(" author:any +")
	}
	if user == "@me" {
		queryString = queryString + fmt.Sprintf(" where:your repos +")
	} else if user == "" {
		queryString = queryString + fmt.Sprintf(" where:anywhere +")
	} else {
		queryString = queryString + fmt.Sprintf(" where:repos by %s +", user)
	}
	if title != "" {
		queryString = queryString + fmt.Sprintf(" title:match +")
	}
	if body != "" {
		queryString = queryString + fmt.Sprintf(" body:match +")
	}
	if len(labelsList) != 0 {
		queryString = queryString + fmt.Sprintf(" labels:")
		for _, label := range labelsList {
			queryString = queryString + fmt.Sprintf("%s ", label)
		}
	}
	queryString = strings.TrimSuffix(queryString, "+")
	queryString = strings.TrimSpace(queryString)
	t := createTable(colour)
	displayInput(t, queryString)
	t.AppendSeparator()
	t.Render()
}

func createTable(textColour string) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)

	width, _, _ := term.GetSize(0)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, WidthMax: 2 * width / 3},
	})

	if colour, ok := colourMap()[textColour]; ok {
		t.Style().Color.Row = text.Colors{colour}
	} else {
		t.Style().Color.Row = text.Colors{text.FgWhite}
	}
	t.Style().Options.SeparateColumns = false
	t.Style().Box.PaddingLeft = "  "
	t.Style().Box.PaddingRight = "  "
	t.Style().Box.BottomLeft = "╰"
	t.Style().Box.TopLeft = "╭"
	t.Style().Box.TopRight = "╮"
	t.Style().Box.BottomRight = "╯"
	return t
}

func displayInput(t table.Writer, queryString string) {
	t.AppendRow(table.Row{queryString})
}

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

	funcMap["getRepoName"] = func(URL string) string {
		urlSlice := strings.Split(URL, "/")
		return urlSlice[3] + "/" + urlSlice[4]
	}

	return &promptui.SelectTemplates{
		Active:   fmt.Sprintf("\U0001F449 {{ .Title | %s | bold }}  {{ .State | faint }}", colour),
		Inactive: fmt.Sprintf("   {{ .Title | %s }}", colour),
		Selected: fmt.Sprintf(`{{ "✔" | green | bold }} {{ .Title | %s | bold }}`, colour),
		Details: `
	{{ "Repository:" | faint }} 	{{ .URL | getRepoName }}
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
