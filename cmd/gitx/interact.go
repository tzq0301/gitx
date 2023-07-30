package main

import (
	"fmt"

	"github.com/go-git/go-git/v5/plumbing/color"
)

const (
	gitAddAll            = "git add -A"
	gitStatus            = "git status"
	gitCommit            = "git commit"
	gitCommitWithMessage = "git commit -m <msg>"
	gitCommitAmend       = "git commit --amend --no-edit"
	gitPullRebase        = "git pull --rebase"
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
	fmt.Printf("[%s %s] %s\n", branchName, hash[:7], commitMessage)
}

func suggestRebase() {
	fmt.Printf("Suggestion: %sRebase%s from %sremote%s or %smain/master%s\n", color.Magenta, color.Reset, color.Green, color.Reset, color.Green, color.Reset)
}
