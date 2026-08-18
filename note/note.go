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
	// json tags are used to specify how the struct fields should be serialized to JSON.
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Display prints the note's title and content to the console.
func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}

// Save saves the note to a JSON file. The filename is derived from the note's title, with spaces replaced by underscores and converted to lowercase.
func (note Note) Save() error{
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(note.Title) + ".json"

	// Creates a json file from the note object.
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	// creates a file adds the content and sets the permissions
	// Write file can return an error. If not error nil is returned
	return os.WriteFile(fileName, json, 0644)
}

// New creates a new Note instance. It returns an error if the title or content is empty.
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
