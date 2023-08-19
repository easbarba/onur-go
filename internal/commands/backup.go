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
	// "path"
	// "strings"

	"github.com/easbarba/onur/internal/common"
	// "github.com/easbarba/onur/internal/database"
)

var backupFolder string = common.BackupFolder()

// Archive will zip repositories and place $DOWNLOADS/archived
func Backup(list *string, verbose *bool) {
	// projectsList := strings.Split(*list, ",")
	// databases, err := database.All() // verbose
	// if err != nil {
	// 	println(err)
	// }

	if *verbose {
		fmt.Printf("\nArchiving at %s\n", backupFolder)
	}

	// for database, _ := range databases {
	// 	for _, p := range database.Projects {
	// 		for _, m := range projectsList {
	// 			if p.Name == path.Base(m) {
	// 				do(m)
	// 			}
	// 		}
	// 	}
	// }
}

// TODO: mkdir archive folder
// TODO: archive to zip by default
// TODO: store at archive folder
func do(project string) {
	fmt.Println(project)
}
