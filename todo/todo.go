package todo

import "fmt"

// TODO: How can I only pass task id to the TodoList struct?

// ================
// Types
// ================

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

// ================
// Listing elements
// ================

func (t *TodoLists) ListAll() []TodoList {
	for _, list := range t.Lists {
		fmt.Println(list.Name)
	}
	return t.Lists // TODO: Remove later? It is needed in test_todo...
}

func (t *TodoList) ListAll() []Task {
	for _, task := range t.Tasks {
		fmt.Println(task.Description)
	}
	return t.Tasks // TODO: Remove later? It is needed in test_todo...
}

func (t *TodoList) ListCompleted() []Task {
	completed_tasks := []Task{}

	for _, task := range t.Tasks {
		if task.Completed {
			fmt.Printf("- [x] %s", task.Description)
			completed_tasks = append(completed_tasks, task)
		}
	}
	return completed_tasks
}

// Add new todo items
// Mark todos as completed
// Delete todos
// Save todos to a file
// Load todos from a file
// Filter todos (show completed/incomplete)
// Basic CLI interface
// Search functionality
