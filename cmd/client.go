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

	var apiResults map[string]interface{}
	err = client.Get("search/issues?"+query.Encode(), &apiResults)
	if err != nil {
		log.Fatal(err)
	}

	itemsResults := apiResults["items"].([]interface{})

	var issues []issueInfo
	for _, item := range itemsResults {
		if item.(map[string]interface{})["pull_request"] == nil {
			issues = append(issues, issueInfo{
				Title: item.(map[string]interface{})["title"].(string),
				URL:   item.(map[string]interface{})["html_url"].(string),
				State: item.(map[string]interface{})["state"].(string),
			})
		}
	}
	return issues
}
