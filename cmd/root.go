package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// VERSION number: change manually
const VERSION = "0.0.0"

var rootCmd = &cobra.Command{
	Use:   "gh-i",
	Short: "gh-i: search repositories interactively",
	Long:  "gh-i: interactive prompt to search and browse github repositories",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if version {
			fmt.Println(VERSION)
			os.Exit(1)
		}
		state, _ := cmd.Flags().GetString("state")
		title, _ := cmd.Flags().GetString("title")
		body, _ := cmd.Flags().GetString("body")
		user, _ := cmd.Flags().GetString("user")
		author, _ := cmd.Flags().GetString("author")
		labelsList, _ := cmd.Flags().GetStringArray("label")
		colour, _ := cmd.Flags().GetString("colour")

		parsedQuery := parseInput(state, title, body, user, author, labelsList)
		issues := getIssues(parsedQuery)
		PromptList := getSelectionPrompt(issues, colour)

		idx, _, err := PromptList.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(issues[idx].URL)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	var labels []string
	rootCmd.Flags().StringP("state", "s", "open", "open or closed issues: default open")
	rootCmd.Flags().StringP("title", "t", "", "search issue by title")
	rootCmd.Flags().StringP("body", "b", "", "search issue by body")
	rootCmd.Flags().StringP("user", "u", "", "search issue in repositories owned by user")
	rootCmd.Flags().StringP("author", "a", "@me", "search issue created by user: default @me")
	rootCmd.Flags().StringArrayVarP(&labels, "label", "l", []string{}, "search issue by label")
	rootCmd.Flags().StringP("colour", "c", "cyan", "colour of selection prompt")
	rootCmd.Flags().BoolP("version", "V", false, "print current version")
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
gh-i: search your github issues interactively.

Synopsis:
	gh i [search] [flags]

Usage:
	gh i

	if no arguments or flags are given, search the user's
	github issues across the web. If the flag -R is provided,
	narrow down the search to a specific repository only.

Prompt commands:

	arrow keys  : move up and down the list
	/           : toggle fuzzy search
	enter (<CR>): open selected repository in the web browser

Flags:
  -R, --repo    only look for issues in a specific repository
  -c, --colour  change prompt colour
  -V, --version print current version
  -h, --help    show this help page
`
}
