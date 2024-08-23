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
	"path"
	"path/filepath"
	"strings"

	"gitlab.com/easbarba/onur/internal/actions"
	"gitlab.com/easbarba/onur/internal/common"
	"gitlab.com/easbarba/onur/internal/storage"
)

// TODO: After grabbing informations log

// Grab all project by pulling or cloning
// TODO return error
func Grab(configName *string, verbose *bool) {
	// configName := topic
	// var configTopic string
	// endWithDot := false

	// if strings.Contains(*topic, ".") {
	// 	parts := strings.SplitN(*topic, ".", 2)
	// 	configName = &parts[0]
	// 	configTopic = parts[1]
	// 	endWithDot = true
	// }

	allConfigs, err := storage.Multi()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if *verbose {
		settings := common.ReadSettings()
		fmt.Println(fmt.Sprintf(`Settings: [ depth: %d, quiet: %t, single-branch: %t ]`,
			settings.Git.Depth, settings.Git.Quiet, settings.Git.SingleBranch))

		fmt.Print("Configurations: [")
		for _, config := range allConfigs {
			fmt.Printf(" %s ", strings.TrimSuffix(filepath.Base(config.Name), ".json"))
		}
		fmt.Println("]")

	}

	for _, config := range allConfigs {
		if *configName != "" && *configName != config.Name {
			continue
		}

		fmt.Println()
		fmt.Println(fmt.Sprintf(`%s: `, config.Name))

		for key, topic := range config.Topics {
			fmt.Println(fmt.Sprintf(`  + %s`, key))

			for _, project := range topic {
				name := strings.ToLower(project.Name)
				folder := path.Join(common.ProjectsFolder(), config.Name, key, name)

				if project.Branch == "" {
					project.Branch = "master"
				}

				projectInfo(name, project.URL, project.Branch)

				if _, err := os.Stat(path.Join(folder, ".git", "config")); err != nil {
					actions.Klone(folder, project.Name, project.URL, project.Branch)
				} else {
					actions.Pull(folder, project.URL, project.Branch)
				}
			}

			fmt.Println()
		}
	}
}
