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

type Task struct {
	ID        string
	Job       string
	Status    string
	CreatedAt string
	UpdatedAt string
}

var addCmd = &cobra.Command{
	Use:   "add task",
	Short: "adds the given task",
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct number of arguments
		if len(args) != 2 {
			fmt.Println("Please provide 2 arg")
			return
		} else {

			if _, err := os.Stat("output.json"); os.IsNotExist(err) {
				file, err := os.Create("output.json")
				if err != nil {
					fmt.Println("erroe creating file:::", err)
					return
				}
				defer file.Close()
			}

			data, err := ioutil.ReadFile("output.json")
			if err != nil && !os.IsNotExist(err) {
				log.Println("Error reading file:", err)
				return
			}

			var task []Task
			// If the file is not empty, unmarshal the content
			if len(data) > 0 {
				err := json.Unmarshal(data, &task)
				if err != nil {
					log.Println("Error unmarshaling JSON:", err)
					return
				}
			}

			newTask := Task{
				ID:        args[0],
				Job:       args[1],
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			}

			task = append(task, newTask)

			log.Println(task)

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

			fmt.Println("Task added to output.json")

		}

	},
}
