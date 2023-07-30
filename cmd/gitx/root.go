package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/color"
	"github.com/samber/lo"
	"github.com/spf13/cobra"

	"github.com/tzq0301/gitx/internal/gitx"
)

var repo *git.Repository

var rootCmd = &cobra.Command{
	Use: "gitx",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	repo = lo.Must(gitx.CurrentRepository())

	branch := gitx.CurrentBranch(repo)

	fmt.Printf("%s%s%s%s\n", promptCurrentBranch, color.Green, branch, color.Reset)

	rootCmd.AddCommand(
		addCmd,
		commitCmd,
		versionCmd,
	)
}
