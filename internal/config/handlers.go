/*
*  Qas is free software: you can redistribute it and/or modify
*  it under the terms of the GNU General Public License as published by
*  the Free Software Foundation, either version 3 of the License, or
*  (at your option) any later version.

*  Qas is distributed in the hope that it will be useful,
*  but WITHOUT ANY WARRANTY; without even the implied warranty of
*  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*  GNU General Public License for more details.

*  You should have received a copy of the GNU General Public License
*  along with Qas. If not, see <https://www.gnu.org/licenses/>.
 */
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/easbarba/qas/internal/common"
	"github.com/easbarba/qas/internal/domain"
)

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

// TODO: Check for duplicates in configuration files
func ConfigCheckDuplicates() {
	panic("unimplemented")
}

// Bundle configurations as a JSON array
func AllToJson() ([]byte, error) {
	// mapped := make(map[string]domain.Projects)

	configs := All()

	// for _, config := range configs {
	// 	mapped[config.Lang] = config.Projects
	// }

	result, err := json.Marshal(configs)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return result, nil
}

// Parse single configuration file,
// returns either properly parsed config parsed or empty struct.
//
//	TODO: check if the expect syntax is correct TODO: or err.
func TranslateConfig(filepath string, filename string) (domain.Config, error) {
	var projects domain.Projects

	file, err := os.ReadFile(filepath)
	if err != nil {
		return domain.Config{}, err
	}

	err = json.Unmarshal(file, &projects)
	if err != nil {
		errMsg := fmt.Sprintf("Configuration file has incorrect syntax \n%s", err.Error())
		return domain.Config{}, errors.New(errMsg)
	}

	config := domain.Config{
		Lang:     common.FileNameWithoutExtension(filename),
		Projects: projects,
	}

	return config, nil
}

func CheckConfigSyntax() error {
	panic("Not implemented")
}

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

func Append(project domain.Projects) domain.Projects {
	panic("not implemented")
}

func Delete(lang string) error {
	configs := All()

	for _, config := range configs {
		if config.Lang == lang {
			err := RemoveConfig(lang)
			if err != nil {
				return errors.New("Unable to delete config")
			}
		}

		return errors.New("No such a configuration file found!")
	}

	// successfully delete configuration file
	return nil
}

func New(payload io.ReadCloser) ([]byte, error) {
	var config domain.Config
	err := json.NewDecoder(payload).Decode(&config)

	if err != nil {
		return nil, errors.New("jackshit")
	}

	err = writeNewConfig(config)
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	return result, nil
}
