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

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/easbarba/onur/internal/common"
	"github.com/easbarba/onur/internal/domain"
)

// return all files found
func Files() []fs.FileInfo {
	onurFolder := common.Configfolder()

	files, err := ioutil.ReadDir(onurFolder)
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
