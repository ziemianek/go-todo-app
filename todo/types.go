package todo

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
