/*
*  Onur is free software: you can redistribute it and/or modify
*  it under the terms of the GNU General Public License as published by
*  the Free Software Foundation, either version 3 of the License, or
*  (at your option) any later version.

*  Onur is distributed in the hope that it will be useful,
*  but WITHOUT ANY WARRANTY; without even the implied warranty of
*  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*  GNU General Public License for more details.

*  You should have received a copy of the GNU General Public License
*  along with Onur. If not, see <https://www.gnu.org/licenses/>.
 */

package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/easbarba/onur/internal/common"
	"github.com/easbarba/onur/internal/domain"
)

// return all files found
func Files() []string {
	var files []string

	entries, err := os.ReadDir(common.Configfolder())
	if err != nil {
		fmt.Printf("Warning: no configuration file found, exiting!")
		os.Exit(1)
	}

	for _, file := range entries {
		fileAbsPath := path.Join(common.Configfolder(), file.Name())

		if ext := filepath.Ext(fileAbsPath); ext == ".json" { // only json files
			if _, err := os.Stat(fileAbsPath); err == nil || !os.IsNotExist(err) { // check if file exist
				files = append(files, fileAbsPath)
			}
		}
	}

	return files
}

func Count() bool {
	panic("not implement!")
}

func exist() {
	panic("not implemented")
}

func to_path() {
	panic("not implemented")
}

// Write new configuration to a json file
func writeNewConfig(newConfig domain.Config) error {
	configs, err := All()

	// Check if any configuration has already Lang set, and skip it!
	for _, config := range configs {
		if config.Topic == newConfig.Topic {
			return errors.New("Configuration already exist. Skipping!")
		}
	}

	// Write new configuration to file
	file, _ := json.Marshal(newConfig.Projects)

	cfgFolder := common.Configfolder()

	newConfigPath := path.Join(cfgFolder, newConfig.Topic+". json")
	err = os.WriteFile(newConfigPath, file, 0644)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func RemoveConfig(lang string) error {
	configFolder := common.Configfolder()

	configPath := path.Join(configFolder, lang+".json")
	err := os.Remove(configPath)
	if err != nil {
		return err
	}

	return nil
}

// Bundle configurations as a JSON array
func AllToJson() ([]byte, error) {
	// mapped := make(map[string]domain.Projects)

	configs, err := All()
	if err != nil {
		return nil, err
	}

	// for _, config := range configs {
	// 	mapped[config.Lang] = config.Projects
	// }

	result, err := json.Marshal(configs)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return result, nil
}
