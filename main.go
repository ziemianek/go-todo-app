package main

import (
	"fmt"
)

type Task struct {
	ID        int
	Name      string
	Completed bool
}

// function to print a menu for todolist app
func printMenu() {
	fmt.Println("Menu:")
	fmt.Println("1. List all tasks")
	fmt.Println("2. Add new task")
	fmt.Println("3. Mark task as completed")
	fmt.Println("4. Delete task")
	fmt.Println("5. Exit")
}

// function to print tasks
func printTasks(allTasks []Task) {
	for index, task := range allTasks {
		fmt.Printf("%d: %s\n", index, task.Name)
	}
}

// function to add a new task to all tasks
func addTask(allTasks []Task, newTask Task) []Task {
	return append(allTasks, newTask)
}

// function to mark a task as completed
func markTaskAsCompleted(allTasks []Task, taskID int) []Task {
	for index, task := range allTasks {
		if task.ID == taskID {
			allTasks[index].Completed = true
		}
	}
	return allTasks
}

// function to delete a task
func deleteTask(allTasks []Task, taskID int) []Task {
	var newTasks []Task
	for _, task := range allTasks {
		if task.ID != taskID {
			newTasks = append(newTasks, task)
		}
	}
	return newTasks
}

func main() {

	allTasks := []Task{}

	fmt.Println("Welcome to my todo list app!")
	printMenu()

	for { // switch case to handle user input
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			printTasks(allTasks)
		case 2: // add new task
			var newTask Task
			fmt.Println("Enter task name:")
			fmt.Scanln(&newTask.Name)
			newTask.ID = len(allTasks) + 1
			newTask.Completed = false
			allTasks = addTask(allTasks, newTask)
			printTasks(allTasks)
		case 3: // mark task as completed
			var taskID int
			fmt.Println("Enter task ID to mark as completed:")
			fmt.Scanln(&taskID)
			allTasks = markTaskAsCompleted(allTasks, taskID)
			printTasks(allTasks)
		case 4: // delete task
			var taskID int
			fmt.Println("Enter task ID to delete:")
			fmt.Scanln(&taskID)
			allTasks = deleteTask(allTasks, taskID)
			printTasks(allTasks)
		case 5: // exit
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
		printMenu()
	}
}
