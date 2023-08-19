package actions

import (
	"fmt"
	"os"

	"github.com/easbarba/onur/internal/common"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// clone repository if none is found at folder
func Klone(folder, name, url, branch string) {
	settings := common.ReadSettings()
	branchHead := fmt.Sprintf("refs/heads/%s", branch)

	_, err := git.PlainClone(folder, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName(branchHead),
		SingleBranch:  settings.SingleBranch,
		Depth:         settings.Depth,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
