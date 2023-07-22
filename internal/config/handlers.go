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
	"io"

	"github.com/easbarba/onur/internal/domain"
)

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
