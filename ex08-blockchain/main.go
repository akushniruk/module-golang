package main

import (
	b "./blockchain"
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "add":
		b.Add()
	case "list":
		b.List()
	case "mine":
		b.Mine()
	case "help":
		fmt.Print("1) go run main.go add [purchase] [cash]\n2) go run main.go list\n3) go run main.go  mine [difficulty]\n")
	default:
		fmt.Print("1) go run main.go add [purchase] [cash]\n2) go run main.go list\n3) go run main.go mine [difficulty]\n")
	}
}
