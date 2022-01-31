package cmd

import (
	"fmt"
	"log"
	"net/url"

	gh "github.com/cli/go-gh"
)

type issueInfo struct {
	Title string
	URL   string
	State string
}

func getIssues(query url.Values) []issueInfo {
	fmt.Println(query)
	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	var apiResults []map[string]interface{}
	err = client.Get("search/issues?"+query.Encode(), &apiResults)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(apiResults)
	var issues []issueInfo
	for _, item := range apiResults {
		if item["pull_request"] == nil {
			issues = append(issues, issueInfo{
				Title: item["title"].(string),
				URL:   item["html_url"].(string),
				State: item["state"].(string),
			})
		}
	}
	return issues
}
