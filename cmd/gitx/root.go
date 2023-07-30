package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
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

	fmt.Printf("Current Branch: %s\n\n", branch.Name())

	rootCmd.AddCommand(
		addCmd,
		versionCmd,
	)
}
