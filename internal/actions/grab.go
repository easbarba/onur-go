// Qas is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Qas is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Qas. If not, see <https://www.gnu.org/licenses/>.

package actions

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/easbarba/qas/internal/common"
	"github.com/easbarba/qas/internal/config"
)

var s = spinner.New(spinner.CharSets[26], 100*time.Millisecond)

// TODO: After grabbing informations log

// Grab all project by pulling or cloning
// TODO return error
func Grab(verbose *bool) {
	projects := config.All() //verbose

	for _, project := range projects {
		for _, pj := range project.Projects {
			name := strings.ToLower(pj.Name)
			folder := path.Join(common.Home(), project.Lang, name)

			printInfo(name, pj.URL, pj.Branch, verbose)

			if _, err := os.Stat(path.Join(folder, ".git")); err == nil {
				pull(folder, pj.URL, pj.Branch)
			} else {
				clone(folder, pj.Name, pj.URL, pj.Branch)
			}
		}
	}
}

func printInfo(name, url, branch string, verbose *bool) {
	title := color.New(color.FgHiYellow, color.Bold).SprintFunc()
	if *verbose {
		fmt.Print(title("name: "), name, title(" url: "), url, title(" branch: "), branch, "\n")
		return
	}

	fmt.Print(title("name: "), name, "\n")
}

// clone repository if none is found at folder
func clone(folder, name, url, branch string) {
	var quiet *os.File

	singleBranch, quietRs, depth := config.Read()
	if quietRs == true {
		quiet = os.Stdout
	} else {
		quiet = os.Stdin
	}

	if branch == "" {
		branch = "master"
	}

	branch = fmt.Sprintf("refs/heads/%s", branch)

	spin.Start()
	_, err := git.PlainClone(folder, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName(branch),
		Progress:      quiet,
		SingleBranch:  singleBranch,
		Depth:         depth,
	})
	spin.Stop()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// pull repository at url/ and branch in the found folder
func pull(folder, url, branch string) {
	var quiet *os.File

	singleBranch, quietRs, depth := config.Read()
	if quietRs == true {
		quiet = os.Stdout
	} else {
		quiet = os.Stdin
	}

	o, err := git.PlainOpen(folder)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w, err := o.Worktree()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spin.Start()
	w.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.ReferenceName(branch),
		SingleBranch:  singleBranch,
		Depth:         depth,
		Progress:      quiet,
	})
	spin.Stop()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
