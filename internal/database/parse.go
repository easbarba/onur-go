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
func One(filepath string) ([]domain.Projects, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return []domain.Projects{}, err
	}

	projects, err := parse(file)
	if err != nil {
		return []domain.Projects{}, err
	}

	return projects, nil
}

func All() ([]domain.Config, error) {
	var configs []domain.Config

	for _, filepath := range Files() {
		one, err := One(filepath)
		if err != nil {

			return []domain.Config{}, err
		}

		config := domain.Config{
			Topic:    common.FileNameWithoutExtension(filepath),
			Projects: one,
		}

		configs = append(configs, config)
	}

	return configs, nil
}

func parse(file []byte) ([]domain.Projects, error) {
	var projects []domain.Projects

	if err := json.Unmarshal(file, &projects); err != nil {
		errMessage := fmt.Sprintf("Configuration file has incorrect syntax \n%s", err.Error())

		return []domain.Projects{}, errors.New(errMessage)
	}

	return projects, nil
}

// TODO: Check for duplicates in configuration files
func ConfigCheckDuplicates() {
	panic("unimplemented")
}

func CheckConfigSyntax() error {
	panic("Not implemented")
}
