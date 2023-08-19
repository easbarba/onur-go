package actions

import (
	"fmt"
	"os"

	"github.com/easbarba/onur/internal/common"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// Pull repository at url/ and branch in the found folder
func Pull(folder, url, branch string) {
	settings := common.ReadSettings()

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

	w.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.ReferenceName(branch),
		SingleBranch:  settings.SingleBranch,
		Depth:         settings.Depth,
	})
}
