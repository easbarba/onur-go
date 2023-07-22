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
	"path"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/easbarba/onur/internal/common"
	"github.com/easbarba/onur/internal/config"
)

var backupFolder string = common.BackupFolder()
var spin = spinner.New(spinner.CharSets[26], 100*time.Millisecond)

// Archive will zip repositories and place $DOWNLOADS/archived
func Backup(list *string, verbose *bool) {
	configs := config.All() // verbose
	projectsList := strings.Split(*list, ",")

	if *verbose {
		fmt.Printf("\nArchiving at %s\n", backupFolder)
	}

	for _, config := range configs {
		for _, p := range config.Projects {
			for _, m := range projectsList {
				if p.Name == path.Base(m) {
					do(m)
				}
			}
		}
	}
}

// TODO: mkdir archive folder
// TODO: archive to zip by default
// TODO: store at archive folder
func do(project string) {
	spin.Start()
	fmt.Println(project)
	spin.Stop()
}
