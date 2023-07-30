package main

import (
	"github.com/spf13/cobra"

	"github.com/tzq0301/gitx/internal/gitx"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: gitAddAll,
	Run: func(cmd *cobra.Command, args []string) {
		executing(gitAddAll)
		gitx.AddAll(repo)
	},
}
