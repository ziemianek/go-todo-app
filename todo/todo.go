package todo

import (
	"fmt"
)

func (t *TodoLists) GetListById(listId int) *TodoList {
	for i := range t.Lists {
		if t.Lists[i].ID == listId {
			return &t.Lists[i]
		}
	}
	return nil
}

// ================
// Listing elements
// ================

func (t *TodoLists) ListAllLists() []TodoList {
	for _, list := range t.Lists {
		fmt.Printf("%d. %s\n", list.ID, list.Name)
	}
	return t.Lists // TODO: Remove later? It is needed in test_todo...
}

func (t *TodoLists) ListAllTasksFromList(listId int) []Task {
	list := t.GetListById(listId)
	for _, task := range list.Tasks {
		fmt.Printf("%d. %s\n", task.ID, task.Description)
	}
	return list.Tasks // TODO: Remove later? It is needed in test_todo...
}

func (t *TodoLists) ListCompleted(listId int) []Task {
	completed_tasks := []Task{}

	list := t.GetListById(listId)

	for _, task := range list.Tasks {
		if task.Completed {
			fmt.Printf("- [x] %s\n", task.Description)
			completed_tasks = append(completed_tasks, task)
		}
	}
	return completed_tasks // TODO: Remove later? It is needed in test_todo...
}

// Add new todo items
func (t *TodoLists) AddTask(listId int, description string) (bool, error) {
	list := t.GetListById(listId)
	task := Task{
		ID:          len(list.Tasks) + 1,
		Description: description,
		Completed:   false,
	}
	list.Tasks = append(list.Tasks, task)

	err := WriteJSON(Filename, t)
	if err != nil {
		panic(err)
	}
	return true, nil
}

// Mark todos as completed
// Delete todos
// Save todos to a file
// Load todos from a file
// Filter todos (show completed/incomplete)
// Basic CLI interface
// Search functionality
