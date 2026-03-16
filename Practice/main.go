package main

import (
	"bufio"
	"fmt"
	"os"
	"practice/demos/note"
	"strings"
)

func main() {
	title, content := getData()
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}
	// fmt.Printf("Title : %s\nContent : %s\nCreated At : %s", newNote.Title, newNote.Content, newNote.CreatedAt)
	newNote.Display()
	err = newNote.Save()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Note saved successfully")

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
