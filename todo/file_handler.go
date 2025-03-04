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
