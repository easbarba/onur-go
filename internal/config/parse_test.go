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

package config

import (
	"path"
	"testing"

	"github.com/easbarba/onur/internal/common"
)

func TestParseConfig(t *testing.T) {
	miscFile := path.Join(common.Configfolder(), "misc.json")
	fileParsed, _ := ParseConfig(miscFile, "misc")

	got := fileParsed.Lang
	expected := "misc"

	if got != expected {
		t.Errorf("Lang attribute does not match, got: %s, expecting %s.", got, expected)
	}

	got = fileParsed.Projects[0].Name
	expected = "awesomewm"

	if got != expected {
		t.Errorf("Name incorrect, got: %s, expecting: %s.", got, expected)
	}
}
