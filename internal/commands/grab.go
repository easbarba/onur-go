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
	projects, err := database.All()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	for _, project := range projects {
		fmt.Println()
		fmt.Println(project.Topic)
		fmt.Println()

		for _, pj := range project.Projects {
			name := strings.ToLower(pj.Name)
			folder := path.Join(common.ProjectsFolder(), project.Topic, name)

			if pj.Branch == "" {
				pj.Branch = "master"
			}

			printInfo(name, pj.URL, pj.Branch, verbose)

			if _, err := os.Stat(path.Join(folder, ".git")); err != nil {
				actions.Klone(folder, pj.Name, pj.URL, pj.Branch)
			} else {
				actions.Pull(folder, pj.URL, pj.Branch)
			}
		}
	}
}

func printInfo(name, url, branch string, verbose *bool) {
	message := name

	if *verbose {
		message += " - " + url + " - " + branch
	}

	fmt.Println(message)
}
