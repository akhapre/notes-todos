package main

import (
	"bufio"
	"fmt"
	"go.ashish.khapre.io/notes-todos/notes"
	"go.ashish.khapre.io/notes-todos/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}
type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()
	text := getTodoData()

	note, err := notes.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(text)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Any value allowed
func getType(data interface{}) string {
	switch data.(type) {
	case notes.Note:
		return "Note"
	default:
		return "Todo"

	}
}
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getTodoData() string {
	content := getUserInput("Todo Text:")
	return content
}
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	//fmt.Scanln(&value)
	value, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return ""
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return value
}
func outputData(data outputtable) error {
	display(data)
	return save(data)

}
func display(data displayer) {
	data.Display()
}

func save(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Print("Error while saving ", getType(data), " : ", err)
		return err
	}
	fmt.Print("Saved ", getType(data), ".")
	return nil
}
