package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/edercarloscosta/clilab/data"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a quick note",
	Long:  `Delete a quick note by Id`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteNote()
	},
}

func init() {
	noteCmd.AddCommand(deleteCmd)
}

func deleteNote() {
	notePromptContent := promptContent{
		"Provide a note ID.",
		"Provide a note ID for delete",
	}
	noteId := promptInput(notePromptContent)

	msgPromptContent := promptContent{
		"Delete agreement.",
		"Are you sure you want to delete? Confirm with Y(Yes) or N(No)",
	}
	confirmation := strings.ToLower(promptInput(msgPromptContent))
	
	if confirmation != "y" {
		fmt.Println("Deletion canceled.")
		return
	}

	id, err := strconv.Atoi(noteId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	data.DeleteById(id)
}
