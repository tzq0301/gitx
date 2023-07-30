package main

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/samber/lo"
	"github.com/spf13/cobra"

	"github.com/tzq0301/gitx/internal/gitx"
)

var (
	commitMessage string
	amend         bool
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: gitCommit,
	Run: func(cmd *cobra.Command, args []string) {
		{
			executing(gofmt)
			gitx.GoFmt()
		}

		{
			executing(gitStatus)
			status := gitx.Status(repo)

			if status.IsClean() {
				fmt.Println("Clean.")
				return
			}

			fmt.Println(status.String())
		}

		{
			var needExecGitAddAll bool
			prompt := &survey.Confirm{
				Message: askExec(gitAddAll),
				Default: true,
			}
			lo.Must0(survey.AskOne(prompt, &needExecGitAddAll))

			if needExecGitAddAll {
				gitx.AddAll(repo)
			}
		}

		{
			if len(commitMessage) != 0 {
				executing(gitCommitWithMessage)
				hash := gitx.Commit(repo, commitMessage)
				showCommitResult(gitx.CurrentBranch(repo), hash.String(), commitMessage)
			} else if amend {
				executing(gitCommitAmend)
				gitx.Amend()
			} else {
				log.Fatalln()
			}
		}

		println()
		suggestRebase()
	},
}

func init() {
	commitCmd.Flags().StringVarP(&commitMessage, "message", "m", "", "")
	commitCmd.Flags().BoolVarP(&amend, "amend", "", false, "")
}
