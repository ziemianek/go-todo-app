package todo

import (
	"testing"
)

func TestTodoListsList_All(t *testing.T) {
	// Initialize a TodoLists with some TodoList items
	todoLists := TodoLists{
		Lists: []TodoList{
			{ID: 1, Name: "List 1", Tasks: []Task{}},
			{ID: 2, Name: "List 2", Tasks: []Task{}},
		},
	}

	// Call the List() method
	lists := todoLists.ListAll()

	// Check the length of the returned list
	if len(lists) != 2 {
		t.Errorf("expected 2 lists, got %d", len(lists))
	}

	// Check the content of the returned list
	expectedNames := []string{"List 1", "List 2"}
	for i, list := range lists {
		if list.Name != expectedNames[i] {
			t.Errorf("expected name %s, got %s", expectedNames[i], list.Name)
		}
	}
}

func TestTodoListListAll(t *testing.T) {
	todoList := TodoList{
		ID:   1,
		Name: "List 1",
		Tasks: []Task{
			{ID: 1, Description: "Task 1", Completed: false},
			{ID: 2, Description: "Task 2", Completed: true},
		},
	}

	tasks := todoList.ListAll()

	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}

	expectedDescriptions := []string{"Task 1", "Task 2"}
	for i, task := range tasks {
		if task.Description != expectedDescriptions[i] {
			t.Errorf("expected description %s, got %s", expectedDescriptions[i], task.Description)
		}
	}
}

func TestTodoListListCompleted(t *testing.T) {
	todoList := TodoList{
		ID:   1,
		Name: "List 1",
		Tasks: []Task{
			{ID: 1, Description: "Task 1", Completed: false},
			{ID: 2, Description: "Task 2", Completed: true},
		},
	}

	completed_tasks := todoList.ListCompleted()

	if len(completed_tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(completed_tasks))
	}

	expectedDescriptions := []string{"Task 2"}
	for i, task := range completed_tasks {
		if task.Description != expectedDescriptions[i] {
			t.Errorf("expected description %s, got %s", expectedDescriptions[i], task.Description)
		}
	}

}
