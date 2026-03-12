package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes-app/note"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter note title: ")
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text())
	fmt.Print("Enter note content: ")
	scanner.Scan()
	content := strings.TrimSpace(scanner.Text())

	n := note.New(title, content)
	txtFile, err := n.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	jsonFile, err := n.SaveAsJSON()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Printf("Note saved to %s and %s\n", txtFile, jsonFile)
}
