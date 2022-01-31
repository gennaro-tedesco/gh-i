package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	inputValues := map[string]interface{}{
		"state":      "open",
		"title":      "support",
		"body":       "body text",
		"user":       "@me",
		"me":         true,
		"labelsList": []string{"bug", "fix"},
	}

	trueString := "type:issue state:open support in:title body text in:body user:@me author:@me label:bug label:fix"
	parsedString := parseInput(inputValues["state"].(string), inputValues["title"].(string), inputValues["body"].(string), inputValues["user"].(string), inputValues["me"].(bool), inputValues["labelsList"].([]string))

	assert.Equal(t, trueString, parsedString["q"][0])
}
