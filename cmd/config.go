package cmd

import (
	"os"
)

/*
PURPOSE:
Handles saving and loading persistent CLI state.

We store current directory in a file:
.clitoolconfig
*/

var configFileName = ".clitoolconfig"

// saveCurrentDir writes the directory to file
func saveCurrentDir(dir string) error {
	return os.WriteFile(configFileName, []byte(dir), 0644)
}

// loadCurrentDir reads directory from file
func loadCurrentDir() string {
	data, err := os.ReadFile(configFileName)
	if err != nil {
		return "." // default directory
	}
	return string(data)
}