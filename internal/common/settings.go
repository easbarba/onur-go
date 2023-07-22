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

type base struct {
	singlebranch bool
	depth        int
}

func ReadSettings() (bool, int) {
	settings_location := path.Join(Configfolder(), "settings.toml")
	if _, err := os.Stat(settings_location); err != nil {
		fmt.Print("no configuration file found", err)
		return true, 1
	}

	var conf base
	_, err := toml.DecodeFile(settings_location, &conf)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		fmt.Println("")
		os.Exit(1)
	}

	return conf.singlebranch, conf.depth
}
