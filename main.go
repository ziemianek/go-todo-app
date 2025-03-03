package main

import (
	"github.com/ziemianek/go-todo-app/todo"
)

func main() {
	// Example data for Task
	task1 := todo.Task{
		ID:          1,
		Description: "Task 1 description",
		Completed:   false,
	}

	task2 := todo.Task{
		ID:          2,
		Description: "Task 2 description",
		Completed:   true,
	}

	// Example data for TodoList
	todoList1 := todo.TodoList{
		ID:       1,
		Name:     "Todo List 1",
		Tasks:    []todo.Task{task1, task2},
		Filename: "todolist1.json",
	}

	todoList2 := todo.TodoList{
		ID:       2,
		Name:     "Todo List 2",
		Tasks:    []todo.Task{},
		Filename: "todolist2.json",
	}

	// Example data for TodoLists
	todoLists := todo.TodoLists{
		Lists: []todo.TodoList{todoList1, todoList2},
	}

	todoLists.ListAll()

}
