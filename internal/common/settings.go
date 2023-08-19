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

	"github.com/BurntSushi/toml"
)

type Settings struct {
	SingleBranch bool `toml:"single_branch"`
	Quiet        bool `toml:"quiet"`
	Depth        int  `toml:"depth"`
}

func ReadSettings() Settings {
	settingsFile := path.Join(Configfolder(), "settings.toml")
	if _, err := os.Stat(settingsFile); err != nil {
		fmt.Print("no configuration file found", err)
		return Settings{true, true, 1}
	}

	var settings Settings
	_, err := toml.DecodeFile(settingsFile, &settings)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		fmt.Println("")

		os.Exit(1)
	}

	return settings
}
