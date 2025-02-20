package note

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

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("empty input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) Print() {
	fmt.Println("---------")
	fmt.Println("Title: ", note.Title)
	fmt.Println("Content: ", note.Content)
	fmt.Println("Creation date: ", note.CreatedAt.String())
	fmt.Println("---------")
}

func (note *Note) Save() error {
	var fileName = strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	var jsonData, err = json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}
