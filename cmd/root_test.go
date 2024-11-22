package cmd

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionFlag(t *testing.T) {
	got := new(bytes.Buffer)
	rootCmd.SetOut(got)
	rootCmd.SetErr(got)
	rootCmd.SetArgs([]string{"--version"})
	err := rootCmd.Execute()
	if assert.NoError(t, err) {
		assert.Regexp(t, regexp.MustCompile(`^\d+\.\d+\.\d+(-\w+)?\n`), got.String())
	}
}
