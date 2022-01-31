<h2 align="center">
  <a href="#" onclick="return false;">
    <img alt="PR" src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat"/>
  </a>
  <a href="https://golang.org/">
    <img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=flat&logo=go&logoColor=white"/>
  </a>
  <a href="https://github.com/gennaro-tedesco/gh-i/releases">
    <img alt="releases" src="https://img.shields.io/github/release/gennaro-tedesco/gh-i"/>
  </a>
</h2>

<h4 align="center">search your github issues interactively</h4>
<h3 align="center">
  <a href="#Installation">Installation</a> •
  <a href="#Usage">Usage</a> •
  <a href="#Feedback">Feedback</a>
</h3>

Search GitHub issues interactively from the command line. Where did you open that bug report three weeks ago? And how many feature requests are still open in your organisation 🤔?

...well say no more:

<img alt="example_image" src="https://user-images.githubusercontent.com/15387611/151801136-c765eca3-c453-453a-ad6b-469ba2e2a454.png">


## Installation
```
gh extension install gennaro-tedesco/gh-i
```
This being a `gh` extension, you of course need [gh cli](https://github.com/cli/cli) as prerequisite.

## Usage
Get started!
```
gh i
```

![demo](https://user-images.githubusercontent.com/15387611/151810424-38095e48-84b7-4a75-8a2e-fa0fb47eaf6f.gif)

Without any flags `gh i` shows all the issues created by yourself in order of last update: this is in the vast majority of cases what you are after, is it not? To refine the search, however, the following flags are availabe
```
gh i [flag]
```
takes one of the following arguments or flags

| flags        | description                                      | example
|:------------ |:------------------------------------------------ |:--------
| --me         | only show issues created by yourself             | gh i --me=false<br>default to true
| -s, --state  | search issues by state (open, closed)            | gh i -s closed<br>default to none, namely both
| -t, --title  | search for issue title                           | gh i -t bug-fix
| -b, --body   | search in issue body                             | gh i -b "not working"
| -u, --user   | search in repos owned by user only               | gh i --me -u @me
| -l, --label  | search for issues by label                       | gh i -l bug -l fix (AND)<br>gh i -l bug,fix (OR)
| -c, --colour | change colour of the prompt                      | gh i -c magenta
| -h, --help   | show the help page                               | gh i -h
| -o, --output | print the output to console<br>default to false, namely open in the browser | gh i -u @me -o | xargs -n1 gh issue view
| -V, --version| print the current version                        | gh i -V

`gh-i` provides the user with visual output of the selected query in human readable format (according to the list of chosen flags):
```
$ gh i -s open -u @me --me=false
╭──────────────────────────────────────────────╮
│  state:open + author:any + where:your repos  │
╰──────────────────────────────────────────────╯
...
```
as well as inline overlay of issue status when browsing through the selection list.

The prompt accepts the following navigation commands:

| key           | description
|:------------- |:-----------------------------------
| arrow keys    | browse results list
| `/`           | toggle search in results list
| `enter (<CR>)`| open issue in browser or return its URL as output (if `-o`)

### Execute commands
`gh-i` must be intended as a filter, to browse the issues you created; as such, the best and most flexible way to execute commands with the results is to pipe it into and from `stdin/stdout`. This said, since in most cases one just wants to view and open the corresponding issue, we default to this action, namely upon selection the issue is opened in the web browser; to override this behaviour and return the output instead, use the `-o` flag.

Check the [Wiki](https://github.com/gennaro-tedesco/gh-i/wiki) for more example and the most common use cases!

## Feedback
If you find this application useful consider awarding it a ⭐, it is a great way to give feedback! Otherwise, any additional suggestions or merge request is warmly welcome!
