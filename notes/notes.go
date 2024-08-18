package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note *Note) Display() {
	fmt.Printf("Title: %v, Content: %v, CreatedAt: %v", note.Title, note.Content, note.CreatedAt)
}

func (note *Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonBytes, err := json.Marshal(note)
	if err != nil {
		fmt.Printf("Error : ", err)
		return err
	}
	return os.WriteFile(fileName, jsonBytes, 0644)
}

func New(title, content string) (*Note, error) {
	if content == "" || title == "" {
		return &Note{}, errors.New("title or content cannot be empty")
	}
	return &Note{Title: title, Content: content, CreatedAt: time.Now()}, nil
}
