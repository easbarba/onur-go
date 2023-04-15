/*
*  Aba is free software: you can redistribute it and/or modify
*  it under the terms of the GNU General Public License as published by
*  the Free Software Foundation, either version 3 of the License, or
*  (at your option) any later version.

*  Aba is distributed in the hope that it will be useful,
*  but WITHOUT ANY WARRANTY; without even the implied warranty of
*  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*  GNU General Public License for more details.

*  You should have received a copy of the GNU General Public License
*  along with Aba. If not, see <https://www.gnu.org/licenses/>.
 */

package common

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/ini.v1"
)

func ReadSettings() (bool, bool, int) {
	settings := path.Join(Configfolder(), "settings.ini")

	cfg, err := ini.Load(settings)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	singleBranch := cfg.Section("common").Key("single-branch").MustBool(false)
	quiet := cfg.Section("common").Key("quiet").MustBool(false)
	depth := cfg.Section("common").Key("depth").MustInt(1)

	return singleBranch, quiet, depth
}
