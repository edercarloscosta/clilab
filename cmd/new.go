package cmd

import (
	"fmt"

	"github.com/edercarloscosta/clilab/data"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new quick note",
	Long:  `Creates a new quick note about your daily activity or work`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

func init() {
	noteCmd.AddCommand(newCmd)
}

func createNewNote() {
	notePromptContent := promptContent{
		"Provide a note.",
		"What message would you like to provide?",
	}
	note := promptInput(notePromptContent)

	categoryPromptContent := promptContent{
		"Provide a category for this note.",
		fmt.Sprintf("Which category would you like to set the note %s?", note),
	}
	category := promptSelect(categoryPromptContent)

	data.Add(note, category)
}
