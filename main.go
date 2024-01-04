package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	command := os.Args[1]

	switch command {
	case "push":
		note := os.Args[2]
		push(note)
		break
	case "pop":
		pop()
		break
	case "clear":
		clear()
		break
	case "show":
		show()
		break
	default:
		os.Exit(-1)

	}
}

func push(note string) {
	appendFile(note)
	fmt.Println("Note updated")
}

func pop() {
	fmt.Println("pop")
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

	fmt.Println("Cleared all notes")
}

func appendFile(note string) {
	file, err := os.OpenFile("./stack.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("-> " + note + "\n")

}
