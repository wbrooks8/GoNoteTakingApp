package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs-practice/note"
)

// main is the entry point of the program. It collects user input for a note's title and content, creates a new Note instance, displays it, and saves it to a JSON file.
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

// getNoteData prompts the user for a note's title and content, returning both as strings.
func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

// getUserInput prompts the user with the provided prompt string and reads a line of input from the command line, returning it as a string.
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
