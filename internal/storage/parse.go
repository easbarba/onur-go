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

package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"gitlab.com/easbarba/onur/internal/common"
	"gitlab.com/easbarba/onur/internal/domain"
)

// Parse single configuration file,
// returns either properly parsed config parsed or empty struct.
func Single(filepath string) (domain.Config, error) {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		return domain.Config{}, err
	}

	topics, err := parse(fileContent)
	if err != nil {
		return domain.Config{}, err
	}

	config := domain.Config{
		Name:   common.FileNameWithoutExtension(filepath),
		Topics: topics,
	}

	return config, nil
}

// returns multiple configurations
func Multi() ([]domain.Config, error) {
	var configs []domain.Config

	for _, filepath := range Files() {
		singleConfig, err := Single(filepath)
		if err != nil {
			return []domain.Config{}, err
		}

		configs = append(configs, singleConfig)
	}

	return configs, nil
}

// TODO: check if the expect syntax is correct TODO: or err.
func parse(file []byte) (domain.Topic, error) {
	var configTopics domain.Topic

	if err := json.Unmarshal(file, &configTopics); err != nil {
		errMessage := fmt.Sprintf("Configuration file has incorrect syntax \n%s", err.Error())

		return domain.Topic{}, errors.New(errMessage)
	}

	return configTopics, nil
}

// TODO: Check for duplicates in configuration files
func ConfigCheckDuplicates() {
	panic("unimplemented")
}

func CheckConfigSyntax() error {
	panic("Not implemented")
}
