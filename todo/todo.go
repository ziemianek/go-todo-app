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

func (t *TodoLists) ListNotCompleted(listId int) error {
	list := t.GetListById(listId)

	for _, task := range list.Tasks {
		if !task.Completed {
			fmt.Printf("- [ ] %s\n", task.Description)
		}
	}
	return nil
}

// Add new todo items
func (t *TodoLists) AddTask(listId int, description string) error {
	list := t.GetListById(listId)
	task := Task{
		ID:          len(list.Tasks) + 1,
		Description: description,
		Completed:   false,
	}
	list.Tasks = append(list.Tasks, task)

	err := WriteJSON(Filename, t)
	if err != nil {
		return err
	}
	return nil
}

// Delete task
func (t *TodoLists) DeleteTask(listId int, taskId int) error {
	list := t.GetListById(listId)

	for i := range list.Tasks {
		if list.Tasks[i].ID == taskId {
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
		}
	}

	err := WriteJSON(Filename, t)
	if err != nil {
		return err
	}
	return nil
}

// Mark todos as completed
func (t *TodoLists) MarkAsCompleted(listId int, taskId int) error {
	list := t.GetListById(listId)

	for i := range list.Tasks {
		if list.Tasks[i].ID == taskId {
			list.Tasks[i].Completed = true
		}
	}

	err := WriteJSON(Filename, t)
	if err != nil {
		return err
	}
	return nil
}
