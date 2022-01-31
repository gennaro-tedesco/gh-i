package cmd

import (
	"fmt"
	"os"

	gh "github.com/cli/go-gh"
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
		me, _ := cmd.Flags().GetBool("me")
		labelsList, _ := cmd.Flags().GetStringArray("label")
		colour, _ := cmd.Flags().GetString("colour")
		output, _ := cmd.Flags().GetBool("output")

		if !output {
			explainInput(state, title, body, user, me, labelsList, colour)
		}
		parsedQuery := parseInput(state, title, body, user, me, labelsList)
		issues := getIssues(parsedQuery)
		PromptList := getSelectionPrompt(issues, colour)

		idx, _, err := PromptList.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if output {
			fmt.Println(issues[idx].URL)
		} else {
			args := []string{"issue", "view", issues[idx].URL, "-w"}
			_, _, err := gh.Exec(args...)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	var labels []string
	rootCmd.Flags().Bool("me", true, "search issues you created")
	rootCmd.Flags().StringP("state", "s", "", "search by open or closed issues")
	rootCmd.Flags().StringP("title", "t", "", "search issue by title")
	rootCmd.Flags().StringP("body", "b", "", "search issue by body")
	rootCmd.Flags().StringP("user", "u", "", "search issue in repositories owned by user")
	rootCmd.Flags().StringArrayVarP(&labels, "label", "l", []string{}, "search issue by label")
	rootCmd.Flags().StringP("colour", "c", "cyan", "colour of selection prompt")
	rootCmd.Flags().BoolP("output", "o", false, "return issue html url to stdout instead of opening it")
	rootCmd.Flags().BoolP("version", "V", false, "print current version")
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
gh-i: search your github issues interactively.

Synopsis:
	gh i [flags]

Usage:
	gh i

	if no arguments or flags are given, search the user's
	github issues across the web.
    After inputting the set of desired flags, the users gets
    visual confirmation of the search query in human readable format.

Prompt commands:

	arrow keys  : move up and down the list
	/           : toggle fuzzy search
	enter (<CR>): open selected repository in the web browser

Flags:
  --me          boolean, only show issues created by yourself?
  -s, --state   search issue by state: open or closed.
                Defaults to nothing, namely both
  -t, --title   search for issue titles
  -b, --body    search in issue body
  -u, --user    search in repositories owned by specified user
  -l, --label   match specific issue labels
                comma separated --> OR (-l="bug,fix")
                many --> AND (-l=bug -l=fix)
  -c, --colour  change prompt colour
  -o, --output  boolean, whether to return the output to console
                (if you want to pipe it into something else).
                Default to false, namely open the issue in the browser
  -V, --version print current version
  -h, --help    show this help page

Examples:

   # search your latest opened issues anywhere
   gh i -s open

   # search all issues in your own repositories
   gh i --me=false -u=@me

   # search the feature you requested long time ago
   gh i -l="feature,feature_request"

   # search by title and body
   gh i -l="bug" -t="upgrade" -b="new version breaks"
`
}
