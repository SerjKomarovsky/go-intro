package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("empty input")
	}
	return Todo{
		text,
	}, nil
}

func (note *Todo) Print() {
	fmt.Println("---------")
	fmt.Println("Text: ", note.Text)
	fmt.Println("---------")
}

func (note *Todo) Save() error {
	var fileName = strings.ReplaceAll(note.Text, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	var jsonData, err = json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}
