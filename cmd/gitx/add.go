package main

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"

	"github.com/tzq0301/gitx/internal/gitx"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "git add -A",
	Run: func(cmd *cobra.Command, args []string) {
		lo.Must0(gitx.AddAll(repo))
	},
}
