package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var marktaskCmd = &cobra.Command{

	Use:   "mark-in-progress task",
	Short: "mark-in-progress the existing task",
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

		for i, t := range task {

			id := t.ID
			log.Println("id::::::", id)
			if id == args[1] {

				t.Status = "in-progress"
				t.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
				task[i] = t

				break // Exit the loop once the update is done

			} else {
				log.Println("no such job id to update the task")
				// return
			}

		}

		log.Println("task marked in progress::::::", task)

		// / Marshal the updated task slice to JSON
		updatedData, err := json.MarshalIndent(task, "", "  ")
		if err != nil {
			log.Println("Error marshaling JSON:", err)
			return
		}

		// Write the updated content back to the file
		err = ioutil.WriteFile("output.json", updatedData, 0644)
		if err != nil {
			log.Println("Error writing to file:", err)
			return
		}

		fmt.Println("task marked in progres in output.json")

	},
}
