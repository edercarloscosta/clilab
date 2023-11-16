package cmd

import (
	"github.com/spf13/cobra"
)

var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A quick note",
	Long:  `A quick note about your daily activity or work`,
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
