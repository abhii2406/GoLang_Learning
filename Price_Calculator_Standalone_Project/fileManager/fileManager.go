package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("fails")
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		file.Close()
		return nil, errors.New("fails")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteJSON(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed")
	}
	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
