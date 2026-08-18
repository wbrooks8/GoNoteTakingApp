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
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v", note.title, note.content)
}

func (note Note) Save() error{
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(note.title)

	// Creates a json file from the note object. 
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	// creates a file adds the content and sets the permissions
	// Write file can return an error. If not error nil is returned
	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
