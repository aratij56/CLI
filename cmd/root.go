package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the root command of the CLI tool
var rootCmd = &cobra.Command{
	Use:   "cli", // Root command name
	Short: "cli is a tool for performing basic mathematical operations",
	Long:  "cli is a tool for performing basic mathematical operations like addition.",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("args::::::", args)
		// If no subcommand is provided or arguments are incorrect
		if len(args) != 3 || args[0] != "add" && args[0] != "update" && args[0] != "delete" {
			log.Println("args[0]", args[0])
			// Print an error message
			if args[0] != "List_all_task" && args[0] != "mark_task" && args[0] != "mark_done" {
				fmt.Println("Error: Invalid command or arguments.")
				fmt.Println("Usage: add task")
				cmd.Help() // Print help message
				os.Exit(1)
			} // Exit with a non-zero status indicating an error
		}

		if args[0] == "add" {

			// Call the addCmd function with arguments
			addCmd.Run(cmd, args[1:])
		} else if args[0] == "update" {

			updateCmd.Run(cmd, args[1:])

		} else if args[0] == "delete" {
			deleteCmd.Run(cmd, args[1:])
		} else if args[0] == "List_all_task" {
			alltaskCmd.Run(cmd, args[0:])

		} else if args[0] == "mark_task" {
			marktaskCmd.Run(cmd, args[0:])
		} else if args[0] == "mark_done" {
			markdoneCmd.Run(cmd, args[0:])
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error occurred while executing CLI: '%s'\n", err)
		os.Exit(1)
	}
}
