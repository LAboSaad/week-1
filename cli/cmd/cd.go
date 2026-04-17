package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/*
PURPOSE:
Simulates changing directory inside the CLI tool.
*/

// shared variable (acts like "current directory")
var currentDir = "."

var cdCmd = &cobra.Command{
	Use:   "cd [directory]",
	Short: "Change directory (persistent)",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]

		info, err := os.Stat(dir)
		if err != nil || !info.IsDir() {
			fmt.Println("Invalid directory:", dir)
			return
		}

		// Save directory persistently
		err = saveCurrentDir(dir)
		if err != nil {
			fmt.Println("Failed to save directory:", err)
			return
		}

		fmt.Println("Changed directory to:", dir)
	},
}