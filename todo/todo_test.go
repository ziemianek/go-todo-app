package todo

import (
	"testing"
)

func TestTodoListsList(t *testing.T) {
	// Initialize a TodoLists with some TodoList items
	todoLists := TodoLists{
		Lists: []TodoList{
			{ID: 1, Name: "List 1", Tasks: []Task{}},
			{ID: 2, Name: "List 2", Tasks: []Task{}},
		},
	}

	// Call the List() method
	lists := todoLists.List()

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
