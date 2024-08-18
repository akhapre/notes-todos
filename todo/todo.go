package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo *Todo) Display() {
	fmt.Printf("Text: %v", todo.Text)
}

func (todo *Todo) Save() error {
	fileName := "todo.json"
	jsonBytes, err := json.Marshal(todo)
	if err != nil {
		fmt.Printf("Error : ", err)
		return err
	}
	return os.WriteFile(fileName, jsonBytes, 0644)
}

func New(content string) (*Todo, error) {
	if content == "" {
		return &Todo{}, errors.New("content cannot be empty")
	}
	return &Todo{Text: content}, nil
}
