package actions

import (
	"fmt"
	"io"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"gitlab.com/easbarba/onur/internal/common"
)

// clone repository if none is found at folder
func Klone(folder, name, url, branch string) {
	settings := common.ReadSettings()
	branchHead := fmt.Sprintf("refs/heads/%s", branch)

	_, err := git.PlainClone(folder, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName(branchHead),
		SingleBranch:  settings.Git.SingleBranch,
		Depth:         settings.Git.Depth,
		Progress:      io.Discard,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
