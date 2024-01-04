package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Push(note string) {
	appendFile("-> " + note)
}

func Pop() {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteContent, _ := io.ReadAll(file)

	content := string(byteContent)

	lines := strings.Split(content, "\n")

	lines = lines[:len(lines)-2]
	Clear()

	for _, line := range lines {
		appendFile(line)
	}

}

func Show() {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteContent, _ := io.ReadAll(file)

	fmt.Println(string(byteContent))
}

func Clear() {
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
