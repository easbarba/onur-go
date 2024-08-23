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

package storage_test

import (
	"path"
	"testing"

	"gitlab.com/easbarba/onur/internal/common"
	"gitlab.com/easbarba/onur/internal/storage"
)

func TestSingleConfig(t *testing.T) {
	miscConfigPath := path.Join(common.Configfolder(), "misc.json")
	singleConfig, err := storage.Single(miscConfigPath)
	if err != nil {
		t.Errorf("Could not parse config")
	}

	got := singleConfig.Name
	expected := "misc"
	if got != expected {
		t.Errorf("Lang attribute does not match, got: %s, expecting %s.", got, expected)
	}

	got = singleConfig.Topics["oss"][0].Name
	expected = "awesomewm"
	if got != expected {
		t.Errorf("Name incorrect, got: %s, expecting: %s.", got, expected)
	}
}

func TestMultiConfig(t *testing.T) {
	multiConfigs, err := storage.Multi()
	if err != nil {
		t.Errorf("Could not parse config")
	}

	got := len(multiConfigs)
	expected := 4
	if got != expected {
		t.Errorf("Expected configs length did not match, got: %d, expecting %d.", got, expected)
	}
}
