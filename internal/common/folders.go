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

package common

import (
	"fmt"
	"os"
	"path"
)

// ProjectsHomeFolder that all projects repositories will be stored at
func ProjectsFolder() string {
	return path.Join(Home(), "Projects")
}

// OnurConfigfolder that config files will be looked up for
func Configfolder() string {
	return path.Join(Home(), ".config", "onur")
}

// Home folder of user
func Home() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error: could not found home!")
		os.Exit(1)
	}

	return home
}

func BackupFolder() string {
	return path.Join(Home(), "Downloads", "archived")
}
