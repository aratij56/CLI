package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var alltaskCmd = &cobra.Command{

	Use:   "list of all tasks",
	Short: "list of all existing task",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("arguments:::::::", args)

		file, err := os.ReadFile("output.json")
		if err != nil {
			log.Println("error while rading file")
		}

		var task []Task
		err = json.Unmarshal(file, &task)
		if err != nil {
			log.Println("error while unmarshalling json")
		}

		log.Println("task:::::::", task)

		fmt.Printf("List of all tasks:\n")

		for _, t := range task {

			fmt.Printf("id: %s, job: \"%s\"\n", t.ID, t.Job)
		}

	},
}
