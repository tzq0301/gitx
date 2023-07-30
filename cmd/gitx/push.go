package main

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/samber/lo"
	"github.com/spf13/cobra"

	"github.com/tzq0301/gitx/internal/gitx"
)

var (
	force bool
)

var pushCmd = &cobra.Command{
	Use: "push",
	Run: func(cmd *cobra.Command, args []string) {
		{
			rebased := false
			prompt := &survey.Confirm{
				Message: askExec(gitPullRebase),
				Default: false,
			}
			lo.Must0(survey.AskOne(prompt, &rebased))

			if !rebased {
				return
			}
		}

		{
			if force {
				executing(gitPushForce)
				gitx.PushForce(repo)
			} else {
				executing(gitPush)
				gitx.Push(repo)
			}
		}
	},
}

func init() {
	pushCmd.Flags().BoolVarP(&force, "force", "f", false, "")
}
