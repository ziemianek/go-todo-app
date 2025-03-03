package todo

import "fmt"

// TODO: How can I only pass task id to the TodoList struct?

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoList struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Tasks    []Task `json:"tasks"`
	Filename string `json:"filename"`
}

type TodoLists struct {
	Lists []TodoList `json:"lists"`
}

// List all existing lists
func (t *TodoLists) List() []TodoList {
	for _, list := range t.Lists {
		fmt.Println(list.Name)
	}
	return t.Lists // TODO: Remove later? It is needed in test_todo...
}

// Add new todo items
// Mark todos as completed
// Delete todos
// Save todos to a file
// Load todos from a file
// Filter todos (show completed/incomplete)
// Basic CLI interface
// Search functionality
