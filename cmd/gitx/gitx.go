package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gitx",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitx~")
	},
}
