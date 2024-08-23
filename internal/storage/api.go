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

package storage

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"gitlab.com/easbarba/onur/internal/common"
)

// Structure structure of Configuration files
// log config files found
type Structure []struct {
	Lang     string `json:"lang"`
	Projects []struct {
		Name   string `json:"name"`
		Branch string `json:"branch"`
		URL    string `json:"url"`
	} `json:"projects"`
}

// TODO: get from url from ENV Variable or cli
const url = "http://localhost:5000/v1/config/list"

// HomeFolder that all projects repositories will be stored at
var HomeFolder string = common.ProjectsFolder()

// All configuration consumed from the onur API
func AllAPI(verbose *bool) Structure {
	onurClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "onur")

	res, getErr := onurClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	config := Structure{}
	jsonErr := json.Unmarshal(body, &config)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return config
}
