package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func removeTaskByID(tasks []Task, id string) []Task {
	var updatedTasks []Task
	for _, t := range tasks {
		if t.ID != id {
			updatedTasks = append(updatedTasks, t)
		}
	}
	return updatedTasks
}

var deleteCmd = &cobra.Command{

	Use:   "delete task",
	Short: "deleting the existing task",
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

		task = removeTaskByID(task, args[0]) // Delete task by ID

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

		fmt.Println("Task deleted  from output.json")

	},
}
