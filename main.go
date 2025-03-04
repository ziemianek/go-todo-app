package main

import (
	"fmt"

	"github.com/ziemianek/go-todo-app/todo"
)

func main() {
	filename := "data/todos.json"

	var todoLists todo.TodoLists
	err := todo.ReadJSON(filename, &todoLists)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	todoLists.ListAll()

}
