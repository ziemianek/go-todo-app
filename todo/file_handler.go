package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

func readTodoListFromJSON(filename string) (*TodoList, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%s does not exist", filename)
		return nil, err
	}
	defer file.Close()

	var todos TodoList

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	return &todos, nil
}
