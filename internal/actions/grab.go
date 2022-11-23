package actions

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/easbarba/qas/internal/config"
)

// TODO: After grabbing log informations

// TODO return error
func Grab(dest string, projects []config.Config) {
	fmt.Printf("Grabbing at:  %s \n", dest)

	for _, project := range projects {
		for _, p := range project.Projects {
			fld := path.Join(dest, project.Lang, p.Name)

			printInfo(fld, p.Name, p.URL, p.Branch)

			if _, err := os.Stat(path.Join(fld, ".git")); err == nil {
				pull(fld, p.URL)
			} else {
				clone(fld, p.Name, p.URL, p.Branch)
			}

		}
	}
}

func printInfo(folder, name, url, branch string) {
	in := `
url: %s
branch: %s
name: %s
folder: %s
`
	fmt.Printf(in, url, branch, name, folder)
}

func clone(folder, name, url, branch string) {
	fmt.Println("status: cloning")
	fmt.Println("")
	branch = fmt.Sprintf("refs/heads/%s", branch)

	_, err := git.PlainClone(folder, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName(branch),
		Progress:      os.Stdout,
		SingleBranch:  true,
		Depth:         1,
	})

	CheckIfError(err)
}

func pull(folder, url string) {
	fmt.Println("status: pulling")
	fmt.Println("")

	o, err := git.PlainOpen(folder)
	CheckIfError(err)

	w, err := o.Worktree()
	CheckIfError(err)

	w.Pull(&git.PullOptions{RemoteName: "origin"})
	CheckIfError(err)
}
