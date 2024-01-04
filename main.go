package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var path string = "C:\\Projects\\KeeperCLI\\stack.txt"

func main() {
	command := os.Args[1]

	switch command {
	case "push":
		if len(os.Args) != 3 {
			panic("No note provided")
		}
		note := os.Args[2]
		push(note)
		fmt.Println("Notes updated")
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
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
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
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteContent, _ := io.ReadAll(file)

	fmt.Println(string(byteContent))
}

func clear() {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func appendFile(note string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(note + "\n")

}
