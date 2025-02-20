/*
 * Note taking app
 * Requirements:
 * Should store notes not only in app memory but also in json file
 * Fields: title, content, createdAt
 */
package main

import (
	"bufio"
	"example/note-app/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Note App!")
	var title, content = getNoteData()

	var userNote, err = note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Print()
	err = userNote.Save()

	if err != nil {
		fmt.Println("saving the note failed.")
		return
	}

	fmt.Println("saving the note succeeded.")
}

func getNoteData() (string, string) {
	var title = getUserInput("Note title: ")
	var content = getUserInput("Note content: ")

	return title, content
}
func getUserInput(prompt string) string {
	fmt.Print(prompt)

	var reader = bufio.NewReader(os.Stdin)

	var text, err = reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
