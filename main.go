package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs-practice/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(promt string) string {
	fmt.Print(promt)
	// for reading more than a single work from cmd line.
	// os.Stdin is the command line.
	reader := bufio.NewReader(os.Stdin)

	// tells reader where to stop reading
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	// removes the \n and \r from the text caputured by reader
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
