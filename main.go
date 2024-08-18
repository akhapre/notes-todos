package main

import (
	"bufio"
	"fmt"
	"go.ashish.khapre.io/notes-todos/notes"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	note, err := notes.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	note.Display()
	err = note.Save()
	if err != nil {
		fmt.Print("Error while saving note : ", err)
	}

}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
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
