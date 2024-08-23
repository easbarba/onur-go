// Onur is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Onur is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Onur. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"
	"strings"

	"gitlab.com/easbarba/onur/internal/storage"
)

func Config(topic, name, url, branch *string, verbose *bool) {
	configName := topic
	var configTopic string
	endWithDot := false

	allConfigs, err := storage.Multi()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if strings.Contains(*topic, ".") {
		parts := strings.SplitN(*topic, ".", 2)
		configName = &parts[0]
		configTopic = parts[1]
		endWithDot = true
	}

	if configTopic != "" {
		for _, config := range allConfigs {
			if config.Name == *configName {
				fmt.Println(fmt.Sprintf(`%s: `, config.Name))

				for key, topic := range config.Topics {
					if key == configTopic {
						fmt.Println(fmt.Sprintf(`  + %s`, key))

						for _, project := range topic {
							name := strings.ToLower(project.Name)
							projectInfo(name, project.URL, project.Branch)
						}
					}
				}
			}
		}
		return
	}

	if endWithDot {
		for _, config := range allConfigs {
			if config.Name == *configName {
				fmt.Print(fmt.Sprintf(`%s: `, config.Name))

				for key, _ := range config.Topics {
					fmt.Print(fmt.Sprintf(` %s `, key))
				}
			}
		}

		return
	}

	for _, config := range allConfigs {
		if config.Name == *configName {
			fmt.Println(fmt.Sprintf(`%s: `, config.Name))

			for key, topic := range config.Topics {
				fmt.Println(fmt.Sprintf(`  + %s`, key))

				for _, project := range topic {
					name := strings.ToLower(project.Name)
					projectInfo(name, project.URL, project.Branch)
				}
			}
		}
	}
}
