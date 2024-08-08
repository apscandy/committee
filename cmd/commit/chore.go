/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package commit

import (
	"fmt"

	"github.com/spf13/cobra"
)

// choreCmd represents the chore command
var choreCmd = &cobra.Command{
	Use:   "chore",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chore called")
	},
}

func init() {
	CommitCmd.AddCommand(choreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// choreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// choreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
