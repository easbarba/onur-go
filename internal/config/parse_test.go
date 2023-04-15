package config

import (
	"path"
	"testing"

	"github.com/easbarba/qas/internal/common"
)

func TestParseConfig(t *testing.T) {
	miscFile := path.Join(common.Configfolder(), "misc.json")
	fileParsed, _ := ParseConfig(miscFile, "misc")

	got := fileParsed.Lang
	expected := "misc"

	if got != expected {
		t.Errorf("Lang attribute does not match, got: %s, expecting %s.", got, expected)
	}

	got = fileParsed.Projects[0].Name
	expected = "awesomewm"

	if got != expected {
		t.Errorf("Name incorrect, got: %s, expecting: %s.", got, expected)
	}
}
