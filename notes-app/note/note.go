package note

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

type Note struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// create a constructor function for Note
func New(title, content string) *Note {
	return &Note{
		ID:        int(time.Now().UnixNano()), // simple unique id from timestamp
		Title:     title,
		Content:   content,
		CreatedAt: time.Now().Format("Jan 2, 2006 at 3:04 PM"),
	}
}

func (n *Note) Save() (string, error) {
	filename := sanitizeFilename(n.Title) + ".txt"
	return filename, os.WriteFile(filename, []byte(n.Content), 0644)
}

func (n *Note) SaveAsJSON() (string, error) {
	filename := sanitizeFilename(n.Title) + ".json"
	data, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return "", err
	}
	return filename, os.WriteFile(filename, data, 0644)
}

func sanitizeFilename(title string) string {
	filename := strings.ToLower(title)
	filename = strings.ReplaceAll(filename, " ", "_")
	filename = strings.ReplaceAll(filename, "/", "-")
	filename = strings.ReplaceAll(filename, "\\", "-")
	return filename
}
