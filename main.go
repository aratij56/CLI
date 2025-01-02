package main

import (
	"cli/cmd" // Ensure the correct import path for the cmd package
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	log.Println("Execute first, then register commands")

	// First, execute any logic you want to run before registering commands (if needed)
	// For example, you could log something or perform some setup here.
	log.Println("Running initial logic...")

	// Now, register the commands

	// Finally, execute the root command
	cmd.Execute() // Runs the root command and processes subcommands like 'add'

}
