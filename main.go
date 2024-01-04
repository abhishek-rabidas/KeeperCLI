package main

import (
	"fmt"
	"os"
)

var path string

func main() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	path = homeDir + "\\notes.txt"

	command := os.Args[1]

	switch command {
	case "push":
		if len(os.Args) != 3 {
			panic("No note provided")
		}
		note := os.Args[2]
		Push(note)
		fmt.Println("Notes updated")
		break
	case "pop":
		Pop()
		fmt.Println("Removed last note")
		break
	case "clear":
		Clear()
		fmt.Println("Cleared all notes")
		break
	case "show":
		Show()
		break
	case "help":
		ShowHelp()
		break
	default:
		fmt.Println("Incorrect command, run [notes help] for seeing all the commands")
		os.Exit(-1)

	}
}
