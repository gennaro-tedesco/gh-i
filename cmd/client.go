package cmd

import (
	"fmt"
	"log"
	"net/url"

	gh "github.com/cli/go-gh"
)

type issueInfo struct {
	Title     string
	URL       string
	UpdatedAt string
}

func getIssues(query url.Values) []issueInfo {
	fmt.Println(query)
	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	var apiResults []map[string]interface{}
	err = client.Get("repos/gennaro-tedesco/gh-s/issues?"+query.Encode(), &apiResults)
	if err != nil {
		log.Fatal(err)
	}

	var issues []issueInfo
	for _, item := range apiResults {
		if item["pull_request"]==nil{
			issues = append(issues, issueInfo{
				Title:     item["title"].(string),
				URL:       item["html_url"].(string),
				UpdatedAt: item["updated_at"].(string),
			})
		}
	}
	return issues
}
