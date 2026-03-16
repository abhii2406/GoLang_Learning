package main

import (
	"bufio"
	"fmt"
	todo "interfaces/learning/Demo"
	"interfaces/learning/note"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type embeddedInterface interface {
	saver
	Display()
}

func main() {
	title, content := getData()
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}
	// fmt.Printf("Title : %s\nContent : %s\nCreated At : %s", newNote.Title, newNote.Content, newNote.CreatedAt)
	// newNote.Display()
	// err = newNote.Save()
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }

	// newNote.Display()
	// err = saveData(newNote)
	err = operations(newNote)
	if err != nil {
		fmt.Print("Fail to save")
		return
	}

	fmt.Println("Note saved successfully")

	text := getUserInput("Text :")

	todo, err := todo.New(text)
	if err != nil {
		fmt.Print(err)
		return
	}
	// fmt.Printf("Title : %s\nContent : %s\nCreated At : %s", newNote.Title, newNote.Content, newNote.CreatedAt)
	// todo.Display()
	// err = todo.Save()
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// todo.Display()
	// err = saveData(todo)
	err = operations(todo)
	if err != nil {
		fmt.Print("Fail to save")
		return
	}
	fmt.Println("Todo saved successfully")

}

func operations(abc embeddedInterface) error {
	abc.Display()
	return saveData(abc)
}
func saveData(abc saver) error {
	err := abc.Save()
	return err
}

func getData() (string, string) {
	title := getUserInput("Title : ")

	content := getUserInput("Content : ")

	return title, content
}

func getUserInput(str string) string {
	fmt.Print(str)

	// var inp string
	// fmt.Scanln(&inp)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	if err != nil {
		return ""
	}

	return text
}
