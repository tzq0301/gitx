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

	branch := lo.Must(gitx.CurrentBranch(repo))

	fmt.Printf("Current Branch: %s %s %s\n", color.Green, branch.Name(), color.Reset)

	rootCmd.AddCommand(
		addCmd,
		versionCmd,
	)
}
