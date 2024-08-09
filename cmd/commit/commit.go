/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commit

import (
	"committee/internal"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long:  `awd`,
	Run: func(cmd *cobra.Command, args []string) {

		options := []string{"fix", "feat", "chore", "refactor", "ci", "build", "docs", "perf"}
		commitType := pterm.DefaultInteractiveSelect.WithOptions(options)
		commitType.DefaultText = "Commit type"
		commitTypeSelected, _ := commitType.Show()

		pterm.Println()

		pterm.DefaultBasicText.Println("Your commit message should complete the following sentence.")
		commitMessage := pterm.DefaultInteractiveTextInput
		commitMessage.DefaultText = "If applied this commit will.."
		commitMessageInput, _ := commitMessage.Show()

		com := internal.NewCommitMessage(commitTypeSelected, commitMessageInput)

		pterm.Println()

		breakingChanges := pterm.DefaultInteractiveConfirm
		breakingChanges.DefaultText = "Does your commit have breaking changes?"
		breakingChangesConfirmed, _ := breakingChanges.Show()
		if breakingChangesConfirmed{
			com.HasBreakingChanges()
		}

		pterm.Println()

		hasScope := pterm.DefaultInteractiveConfirm
		hasScope.DefaultText = "Does your commit have a scope of work?"
		hasScopeConfirmed, _ := hasScope.Show()
		pterm.Println()

		if hasScopeConfirmed {
			scopeMessage := pterm.DefaultInteractiveTextInput
			scopeMessage.DefaultText = "Enter your scope here"
			scopeMessageInput, _ := scopeMessage.Show()

			pterm.Println()
			com.AddScope(scopeMessageInput)
		}

		confirmCommit := pterm.DefaultInteractiveConfirm
		pterm.DefaultBasicText.Println("your commit message will be " + pterm.Red(com.ToSting()))
		confirmCommit.DefaultText = "Are you happy with this commit message?" 
		confirmCommitBool, _ := confirmCommit.Show()
		if confirmCommitBool{
			internal.Commit(com)
		}
		os.Exit(1)
	},
}
