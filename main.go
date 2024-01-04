package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	command := os.Args[1]

	switch command {
	case "push":
		note := os.Args[2]
		if len(note) == 0 {
			panic("No note provided")
		}
		push(note)
		fmt.Println("Note updated")
		break
	case "pop":
		pop()
		fmt.Println("Removed last note")
		break
	case "clear":
		clear()
		fmt.Println("Cleared all notes")
		break
	case "show":
		show()
		break
	default:
		os.Exit(-1)

	}
}

func push(note string) {
	appendFile("-> " + note)
}

func pop() {
	file, err := os.OpenFile("./stack.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteContent, _ := io.ReadAll(file)

	content := string(byteContent)

	lines := strings.Split(content, "\n")

	lines = lines[:len(lines)-2]
	clear()

	for _, line := range lines {
		appendFile(line)
	}

}

func show() {
	file, err := os.OpenFile("./stack.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteContent, _ := io.ReadAll(file)

	fmt.Println(string(byteContent))
}

func clear() {
	file, err := os.OpenFile("./stack.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func appendFile(note string) {
	file, err := os.OpenFile("./stack.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(note + "\n")

}
