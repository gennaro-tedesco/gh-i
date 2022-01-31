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
		filter, _ := cmd.Flags().GetString("filter")
		state, _ := cmd.Flags().GetString("state")
		labelsList, _ := cmd.Flags().GetStringArray("label")
		sort, _ := cmd.Flags().GetString("sort")
		colour, _ := cmd.Flags().GetString("colour")

		parsedQuery := parseInput(filter, state, labelsList, sort)
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
	rootCmd.Flags().StringP("filter", "f", "created", "one of 'assigned,created,mentioned,subscribed,all': default 'createad'")
	rootCmd.Flags().StringArrayVarP(&labels, "label", "l", []string{}, "search issue by label")
	rootCmd.Flags().StringP("state", "S", "open", "one of 'open,closed,all': default 'open'")
	rootCmd.Flags().StringP("sort", "s", "updated", "one of 'created,updated': default 'updated'")
	rootCmd.Flags().StringP("repo", "R", "", "search issues in a specific repository")
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
