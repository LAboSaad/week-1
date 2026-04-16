package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var port int
var verbose bool

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a simple HTTP server",

	Run: func(cmd *cobra.Command, args []string) {

		// Inform user server is starting
		fmt.Println("Server running on port", port)

		// Optional verbose logging
		if verbose {
			fmt.Println("Verbose mode enabled")
		}

		/*
			HTTP ROUTE SETUP

			This defines what happens when user visits:
				http://localhost:PORT/
		*/
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello from CLI server")
		})

		// Convert port number to string format ":3000"
		addr := fmt.Sprintf(":%d", port)

		fmt.Println("Listening on http://localhost" + addr)

		/*
			START SERVER (BLOCKING CALL)

			This line:
			- starts listening for requests
			- keeps program running
			- blocks further execution
		*/
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			fmt.Println("Server failed:", err)
		}
	},
}

func init() {

	// Port flag (short: -p)
	serveCmd.Flags().IntVarP(
		&port,
		"port",
		"p",
		8080,
		"Port to run server on",
	)

	// Verbose flag (short: -v)
	serveCmd.Flags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"Enable verbose logging",
	)
}
