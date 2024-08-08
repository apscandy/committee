/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commit

import (
	"committee/internal"
	"fmt"

	"github.com/spf13/cobra"
)

// featCmd represents the feat command
var featCmd = &cobra.Command{
	Use:   "feat",
	Short: "commit message for a feature",
	Long:  `commit message for a feature`,

	Run: func(cmd *cobra.Command, args []string) {
		com := internal.NewCommitMessage("feat", message)
		internal.Commit(com)
	},
}

func init() {
	featCmd.Flags().StringVarP(&message, "message", "m", "", "")

	if err := featCmd.MarkFlagRequired("message"); err != nil {
		fmt.Println(err.Error())
	}

	CommitCmd.AddCommand(featCmd)
}
