package todo

import (
	"encoding/json"
	"io"
	"os"
)

func ReadJSON(filename string, v interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, v)
}

func WriteJSON(filename string, v interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the data into JSON and write it to the file
	encoder := json.NewEncoder(file)
	return encoder.Encode(v)
}
