package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string
}

func New(str string) (Todo, error) {
	if str == "" {
		return Todo{}, errors.New("Text is Empty")
	}
	return Todo{
		Text: str,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("Text : %v", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
