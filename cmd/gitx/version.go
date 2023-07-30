package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := exec.Command("git", "version").Output()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Print(string(output))
	},
}
