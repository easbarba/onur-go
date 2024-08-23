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

package common_test

import (
	"testing"

	"gitlab.com/easbarba/onur/internal/common"
)

func TestReadSettings(t *testing.T) {
	expectedSingleBranch := true
	settings := common.ReadSettings()

	if settings.Git.SingleBranch != expectedSingleBranch {
		t.Errorf("Expected %t, got %t instead.\n", expectedSingleBranch, settings.Git.SingleBranch)
	}

	expectedQuiet := false
	if settings.Git.Quiet != expectedQuiet {
		t.Errorf("Expected %t, got %t instead.\n", expectedQuiet, settings.Git.Quiet)
	}

	expectedDepth := 3
	if settings.Git.Depth != expectedDepth {
		t.Errorf("Expected %d, got %d instead.\n", expectedDepth, settings.Git.Depth)
	}
}
