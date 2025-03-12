package main

import (
	"fmt"

	"github.com/ziemianek/go-todo-app/todo"
)

func main() {
	var todoLists todo.TodoLists
	err := todo.ReadJSON(todo.Filename, &todoLists)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	listId := 1
	todoLists.ListAllTasksFromList(listId)
	// todoLists.AddTask(listId, "New task")
	// todoLists.ListAllTasksFromList(listId)

}
