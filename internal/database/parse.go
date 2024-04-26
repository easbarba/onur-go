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

	"github.com/easbarba/onur/internal/common"
	"github.com/easbarba/onur/internal/domain"
)

// Parse single configuration file,
// returns either properly parsed config parsed or empty struct.
//
//	TODO: check if the expect syntax is correct TODO: or err.
func One(filepath string) (domain.Topic, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return domain.Topic{}, err
	}

	subtopics, err := parse(file)
	if err != nil {
		return domain.Topic{}, err
	}

	return subtopics, nil
}

func All() ([]domain.Config, error) {
	var configs []domain.Config

	for _, filepath := range Files() {
		one, err := One(filepath)
		if err != nil {
			return []domain.Config{}, err
		}

		config := domain.Config{
			Name:  common.FileNameWithoutExtension(filepath),
			Topic: one,
		}

		configs = append(configs, config)
	}

	return configs, nil
}

func parse(file []byte) (domain.Topic, error) {
	var subtopics domain.Topic

	if err := json.Unmarshal(file, &subtopics); err != nil {
		errMessage := fmt.Sprintf("Configuration file has incorrect syntax \n%s", err.Error())

		return domain.Topic{}, errors.New(errMessage)
	}

	return subtopics, nil
}

// TODO: Check for duplicates in configuration files
func ConfigCheckDuplicates() {
	panic("unimplemented")
}

func CheckConfigSyntax() error {
	panic("Not implemented")
}
