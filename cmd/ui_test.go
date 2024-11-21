package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	type testCase struct {
		inputValues map[string]interface{}
		trueString  string
	}
	testCases := map[string]testCase{
		"skip to set repo if given empty": {
			inputValues: map[string]interface{}{
				"state":      "open",
				"title":      "support",
				"body":       "body text",
				"user":       "@me",
				"repo":       "",
				"me":         true,
				"labelsList": []string{"bug", "fix"},
			},
			trueString: "type:issue state:open support in:title body text in:body user:@me author:@me label:bug label:fix",
		},

		"set repo if specified": {
			inputValues: map[string]interface{}{
				"state":      "open",
				"title":      "support",
				"body":       "body text",
				"user":       "@me",
				"repo":       "gennaro-tedesco/gh-f",
				"me":         true,
				"labelsList": []string{"bug", "fix"},
			},
			trueString: "type:issue state:open support in:title body text in:body user:@me repo:gennaro-tedesco/gh-f author:@me label:bug label:fix",
		},
	}

	for description, testCase := range testCases {
		t.Run(description, func(t *testing.T) {
			inputValues := testCase.inputValues

			parsedString := parseInput(inputValues["state"].(string), inputValues["title"].(string), inputValues["body"].(string), inputValues["user"].(string), inputValues["repo"].(string), inputValues["me"].(bool), inputValues["labelsList"].([]string))
			assert.Equal(t, testCase.trueString, parsedString["q"][0])
		})
	}
}

func TestParseRepo(t *testing.T) {
	type testCase struct {
		input  string
		ok     bool
		want   string
		errmsg string
		envs   map[string]string
	}

	testCases := map[string]testCase{
		"given a valid format": {
			input: "gennaro-tedesco/gh-f",
			ok:    true,
			want:  "gennaro-tedesco/gh-f",
		},
		"given an invalid format": {
			input:  "gh-i",
			ok:     false,
			errmsg: `expected the "OWNER/REPO" format, got "gh-i"`,
		},
		"not given a repo but having only $GH_REPO": {
			input: "",
			ok:    true,
			want:  "",
			envs: map[string]string{
				"GH_REPO": "gennaro-tedesco/gh-f",
			},
		},
		"not given a repo but having both $GH_REPO and $GH_I_PREFER_REPO": {
			input: "",
			ok:    true,
			want:  "gennaro-tedesco/gh-f",
			envs: map[string]string{
				"GH_REPO":          "gennaro-tedesco/gh-f",
				"GH_I_PREFER_REPO": "true",
			},
		},
		// No test case added for "no input and no $GH_REPO" due to gh command respecting current directory.
		// Preferring "gh-f" over "gh-i" for a fixture is rooted in the same reason.
	}

	for description, testCase := range testCases {
		t.Run(description, func(t *testing.T) {
			for name, value := range testCase.envs {
				t.Setenv(name, value)
			}
			got, err := parseRepo(testCase.input)
			ok := err == nil
			if ok != testCase.ok {
				assert.FailNow(t, "given an unexpected result")
			}
			if testCase.ok {
				assert.Equal(t, testCase.want, got)
			} else {
				assert.Equal(t, testCase.errmsg, err.Error())
			}
		})
	}
}
