/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commit

import (
	"committee/internal"
	"fmt"

	"github.com/spf13/cobra"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		com := internal.NewCommitMessage("docs", message)
		internal.Commit(com)
	},
}

func init() {
	docsCmd.Flags().StringVarP(&message, "message", "m", "", "")

	if err := docsCmd.MarkFlagRequired("message"); err != nil {
		fmt.Println(err.Error())
	}

	CommitCmd.AddCommand(docsCmd)
}
