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
	"os"
	"path"
	"path/filepath"

	"github.com/easbarba/qas/internal/common"
	"github.com/easbarba/qas/internal/domain"
)

// return all configurations or none.
func All() []domain.Config {
	var configs []domain.Config

	for _, file := range Files() {
		cfgDir := common.Configfolder()

		p := path.Join(cfgDir, file.Name())
		fileInfo, err := os.Stat(p)

		// ignore broken symbolic link
		if os.IsNotExist(err) {
			continue
		}

		// ignore directories
		if fileInfo.IsDir() {
			continue
		}

		// ignore csv files (legacy)
		if ext := filepath.Ext(p); ext == ".csv" {
			continue
		}

		// ignore ini files (legacy)
		if ext := filepath.Ext(p); ext == ".ini" {
			continue
		}
		// ignore toml files
		if ext := filepath.Ext(p); ext == ".toml" {
			continue
		}

		configed, err := ParseConfig(p, file.Name())
		if err != nil {
			fmt.Print(err)
			continue // just ignore faulty/empty configs
		}

		configs = append(configs, configed)
	}

	return configs
}

func One(lang string) ([]byte, error) {
	configs := All()

	for _, config := range configs {
		if config.Lang == lang {
			cfg, err := json.Marshal(config)
			if err != nil {
				return nil, errors.New("Unable to convert current config to JSON")
			}

			return cfg, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("No configuration with Lang '%s' tag found!", lang))
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
