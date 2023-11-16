package cmd

import (
	"github.com/spf13/cobra"
	"github.com/edercarloscosta/clilab/data"
)	

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all quick notes you've added",
	Long:  `Get all quick notes you've added`,
	Run: func(cmd *cobra.Command, args []string) {
		data.GetAll()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
} 