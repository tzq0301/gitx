package gitx

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/samber/lo"
)

func CurrentRepository() (*git.Repository, error) {
	return git.PlainOpen(".")
}

func CurrentBranch(repo *git.Repository) string {
	return strings.TrimLeft(lo.Must(repo.Head()).Name().String(), "refs/heads/")
}

func CurrentWorktree(repo *git.Repository) *git.Worktree {
	return lo.Must(repo.Worktree())
}

func AddAll(repo *git.Repository) {
	worktree := CurrentWorktree(repo)

	lo.Must0(worktree.AddWithOptions(&git.AddOptions{
		All: true,
	}))
}

func Commit(repo *git.Repository, message string) plumbing.Hash {
	worktree := CurrentWorktree(repo)
	return lo.Must(worktree.Commit(message, &git.CommitOptions{}))
}

func Status(repo *git.Repository) *git.Status {
	worktree := CurrentWorktree(repo)
	status := lo.Must(worktree.Status())
	return &status
}

func Amend() {
	lo.Must0(exec.Command("git", "commit", "--amend", "--no-edit").Run())
}

func Push(repo *git.Repository) {
	err := lo.Must(repo.Remote("origin")).Push(&git.PushOptions{
		Progress: os.Stdout,
	})
	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		return
	}
	lo.Must0(err)
}

func PushForce(repo *git.Repository) {
	err := lo.Must(repo.Remote("origin")).Push(&git.PushOptions{
		Force:    true,
		Progress: os.Stdout,
	})
	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		return
	}
	lo.Must0(err)
}
