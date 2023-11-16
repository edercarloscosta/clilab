package cmd

import (
"github.com/spf13/cobra"
"github.com/edercarloscosta/clilab/data"
)

var initCmd = &cobra.Command{
	Use: "init",
	Short: "Initialize a new quick note ",
	Long: `Initialize the database & table structure for quick notes`,
	Run: func (cmd *cobra.Command, args []string) {
		data.CreateTable()
	} ,
} 

func init(){
	rootCmd.AddCommand(initCmd)
}

