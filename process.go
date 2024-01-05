package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Push(note string) {
	appendFile(note)
}

func Pop() {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("No Notes")
		return
	}
	defer file.Close()

	byteContent, _ := io.ReadAll(file)

	content := string(byteContent)

	lines := strings.Split(content, "\n")

	if len(lines) == 1 {
		fmt.Println("No notes")
		os.Exit(-1)
	}

	lines = lines[:len(lines)-2]

	Clear()

	for _, line := range lines {
		appendFile(line)
	}

}

func Show() {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("No Notes")
		os.Exit(-1)
	}
	defer file.Close()

	byteContent, _ := io.ReadAll(file)

	content := string(byteContent)

	lines := strings.Split(content, "\n")

	if len(lines) == 1 {
		fmt.Println("No notes")
		os.Exit(-1)
	}

	lines = lines[:len(lines)-1]

	ShowList(lines)

	//fmt.Println(string(byteContent))
}

func Clear() {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("No Notes")
		os.Exit(-1)
	}
	defer file.Close()
}

func RewriteNotes(selected string) {
	newNotes := make([]string, 0)

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("No Notes")
		os.Exit(-1)
	}
	defer file.Close()

	byteContent, _ := io.ReadAll(file)

	content := string(byteContent)

	lines := strings.Split(content, "\n")

	if len(lines) == 1 {
		fmt.Println("No notes")
		os.Exit(-1)
	}

	lines = lines[:len(lines)-1]

	for _, note := range lines {
		if note == selected {
			continue
		} else {
			newNotes = append(newNotes, note)
		}
	}

	Clear()

	for _, line := range newNotes {
		appendFile(line)
	}
}

func appendFile(note string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening notes")
		os.Exit(-1)
	}
	defer file.Close()

	file.WriteString(note + "\n")

}

func ShowHelp() {
	fmt.Println("CLI NOTE KEEPING APPLICATION")
	fmt.Println("Developed by : Abhishek Kumar Rabidas")
	fmt.Println("1. push [enter your note] : To add a new note")
	fmt.Println("2. clear : To clear all the notes")
	fmt.Println("3. pop : To remove the last note")
	fmt.Println("4. show : To show all the notes")
}
