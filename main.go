package main

import (
	"fmt"
	"os"
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
	default:
		os.Exit(-1)

	}
}
