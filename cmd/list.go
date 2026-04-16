package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/*

It helps users inspect filesystem contents quickly.
────────────────────────────────────────────────────
*/

/*
listCmd defines the "list" subcommand.

WHY THIS EXISTS:
Cobra uses a command object model where each CLI command
is represented as a struct.

This command will be triggered like:

	./cli-tool list
*/
var listCmd = &cobra.Command{
	Use:   "list", // command name
	Short: "List files in current directory",

	/*
		Run function executes when user runs:
			./cli-tool list

		WHAT IT DOES:
		- Reads current directory
		- Prints all file names
	*/
	Run: func(cmd *cobra.Command, args []string) {

		// Read current directory (".")
		dir := loadCurrentDir()
		files, err := os.ReadDir(dir)
		fmt.Println("Listing directory:", dir)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}

		// Loop through files and print names
		for _, file := range files {
			fmt.Println(file.Name())
		}
	},
}
