package main

import (
	"fmt"

	"github.com/go-git/go-git/v5/plumbing/color"
)

const (
	gitAddAll = "git add -A"
	gitStatus = "git status"
	gitCommit = "git commit"
)

const (
	gofmt = "gofmt -w ."
)

const (
	promptCurrentBranch = "Current Branch: "
	promptExecuting     = "Executing: "
)

func executing(name string) {
	fmt.Printf("%s%s%s%s\n", promptExecuting, color.Blue, name, color.Reset)
}

func askExec(s string) string {
	return fmt.Sprintf("Exec %s%s%s >", color.Blue, s, color.Reset)
}

func showCommitResult(branchName string, hash string, commitMessage string) {
	fmt.Printf("[%s %s] %s\n", branchName, hash, commitMessage)
}
