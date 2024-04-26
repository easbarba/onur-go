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

package commands

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/easbarba/onur/internal/actions"
	"github.com/easbarba/onur/internal/common"
	"github.com/easbarba/onur/internal/database"
)

// TODO: After grabbing informations log

// Grab all project by pulling or cloning
// TODO return error
func Grab(verbose *bool) {
	settings := common.ReadSettings()
	fmt.Println(fmt.Sprintf(`settings { depth: %d, quiet: %t, single-branch: %t }`, settings.Depth, settings.Quiet, settings.SingleBranch))

	allProjects, err := database.All()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	for _, project := range allProjects {
		fmt.Println()
		fmt.Println(fmt.Sprintf(`> %s`, project.Name))

		for key, topic := range project.Topic {
			fmt.Println(fmt.Sprintf(`  + %s`, key))

			for _, projekt := range topic {
				name := strings.ToLower(projekt.Name)
				folder := path.Join(common.ProjectsFolder(), project.Name, key, name)

				if projekt.Branch == "" {
					projekt.Branch = "master"
				}

				projectInfo(name, projekt.URL, projekt.Branch)

				if _, err := os.Stat(path.Join(folder, ".git", "config")); err != nil {
					actions.Klone(folder, projekt.Name, projekt.URL, projekt.Branch)
				} else {
					actions.Pull(folder, projekt.URL, projekt.Branch)
				}
			}

			fmt.Println()
		}
	}
}

func projectInfo(name, url, branch string) {
	message := fmt.Sprintf(`    - %-35s %-75s  %s`, name, url, branch)
	fmt.Println(message)
}
