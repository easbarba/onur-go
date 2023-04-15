package config

import (
	"fmt"
	"os"
	"path"

	"github.com/easbarba/qas/internal/common"
	"gopkg.in/ini.v1"
)

func Read() (bool, bool, int) {
	settings := path.Join(common.Configfolder(), "settings.ini")

	cfg, err := ini.Load(settings)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	singleBranch := cfg.Section("common").Key("single-branch").MustBool(false)
	quiet := cfg.Section("common").Key("quiet").MustBool(false)
	depth := cfg.Section("common").Key("depth").MustInt(1)

	return singleBranch, quiet, depth
}
