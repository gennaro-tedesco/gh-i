<h2 align="center">
  <a href="#" onclick="return false;">
    <img alt="PR" src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat"/>
  </a>
  <a href="https://golang.org/">
    <img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=flat&logo=go&logoColor=white"/>
  </a>
  <a href="https://github.com/gennaro-tedesco/gh-s/releases">
    <img alt="releases" src="https://img.shields.io/github/release/gennaro-tedesco/gh-s"/>
  </a>
</h2>

<h4 align="center">search your github issues interactively</h4>
<h3 align="center">
  <a href="#Installation">Installation</a> •
  <a href="#Usage">Usage</a> •
  <a href="#Feedback">Feedback</a>
</h3>

## Installation
```
# local installation
gh repo clone gennaro-tedesco/gh-i
go build .
gh install .
```
This being a `gh` extension, you of course need [gh cli](https://github.com/cli/cli) as prerequisite.

## Usage
Get started!
```
gh i
```
Without any flags `gh i` shows all the issues created by yourself in order of last update: this is more often than not what you are after, is it not? To refine the search, however, the following flags are availabe
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
| -V, --version| print the current version                        | gh i -V

The prompt accepts the following navigation commands:

| key           | description
|:------------- |:-----------------------------------
| arrow keys    | browse results list
| `/`           | toggle search in results list
| `enter (<CR>)`| open selected repository in web browser
