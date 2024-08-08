/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commit

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	message string
)

// featCmd represents the feat command
var featCmd = &cobra.Command{
	Use:   "feat",
	Short: "commit message for a feature",
	Long: `commit message for a feature`,

	Run: func(cmd *cobra.Command, args []string) {
		var sb strings.Builder
		sb.WriteString("feat: ")
		sb.WriteString(message)

		execCmd := exec.Command("git", "commit", "-m", sb.String())
		stdout, err := execCmd.Output()

		if err != nil {
			fmt.Println("failed to commit double check you have staged your changes")
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(stdout))
	},
}

func init() {
	featCmd.Flags().StringVarP(&message, "message", "m", "", "")

	if err := featCmd.MarkFlagRequired("message"); err != nil{
		fmt.Println(err.Error())
	}

	CommitCmd.AddCommand(featCmd)
}
