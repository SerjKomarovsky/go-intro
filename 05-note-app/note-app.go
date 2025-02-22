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
	"example/note-app/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
    Save() error
}

// type printer interface {
//     Print()
// }

// type outputtable interface {
//     Save() error
//     Print() 
// }
type outputtable interface {
    saver
    Print()
}

func main() {
	fmt.Println("Welcome to Note App!")
	var title, content = getNoteData()
    var todoText = getUserInput("Todo text: ")

    todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

    err = printData(&todo)
    // todo.Print()
    // err = saveData(&todo)

    if err != nil {
        return
    }

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

    err = printData(&userNote)
	// userNote.Print()
	// err = saveData(&userNote)

    if err != nil {
        return
    }
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

func saveData(data saver) error {
    err := data.Save()

	if err != nil {
		fmt.Println("saving the data failed.")
		return err
	}

	fmt.Println("saving the data succeeded.")
    return nil
}

func printData(data outputtable) error {
    data.Print()
    return saveData(data)
}

// example of embedded interface
func printSomething(value any) {
    switch value.(type) {
        case int: 
            fmt.Println("Integer:", value)
        case float64: 
            fmt.Println("Float:", value)
        case string: 
            fmt.Println("String:", value)
    }
}
