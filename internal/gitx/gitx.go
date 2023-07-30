package gitx

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func CurrentRepository() (*git.Repository, error) {
	return git.PlainOpen(".")
}

func CurrentBranch(repo *git.Repository) (*plumbing.Reference, error) {
	return repo.Head()
}

func AddAll(repo *git.Repository) error {
	worktree, err := repo.Worktree()

	if err != nil {
		return err
	}

	err = worktree.AddWithOptions(&git.AddOptions{
		All: true,
	})

	if err != nil {
		return err
	}

	return nil
}
