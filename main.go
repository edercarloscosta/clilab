package main

import (
	"github.com/edercarloscosta/clilab/cmd"
	"github.com/edercarloscosta/clilab/data"
)

func main() {
	data.OpenDb()
	cmd.Execute()	
} 