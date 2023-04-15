package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/easbarba/qas/internal/common"
	"github.com/easbarba/qas/internal/domain"
)

// return all files found
func Files() []fs.FileInfo {
	qasFolder := common.Configfolder()

	files, err := ioutil.ReadDir(qasFolder)
	if err != nil {
		fmt.Printf("Warning: no configuration file found, exiting!")
		os.Exit(1)
	}

	return files
}

// Write new configuration to a json file
func writeNewConfig(newConfig domain.Config) error {
	configs := All()

	// Check if any configuration has already Lang set, and skip it!
	for _, config := range configs {
		if config.Lang == newConfig.Lang {
			return errors.New("Configuration already exist. Skipping!")
		}
	}

	// Write new configuration to file
	file, _ := json.Marshal(newConfig.Projects)

	cfgFolder := common.Configfolder()

	newConfigPath := path.Join(cfgFolder, newConfig.Lang+".json")
	err := os.WriteFile(newConfigPath, file, 0644)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func RemoveConfig(lang string) error {
	cfgFolder := common.Configfolder()

	configPath := path.Join(cfgFolder, lang+".json")
	err := os.Remove(configPath)
	if err != nil {
		return err
	}

	return nil
}
