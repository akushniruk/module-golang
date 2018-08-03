package blockchain

import (
	"fmt"
	"os"
	"strconv"
)

func List() {
	block := ReadFile(Filename)
	for count := 0; count < len(block.Blocks); count++ {
		fmt.Println("Index: " + strconv.Itoa(block.Blocks[count].Index))
		fmt.Println("Time: " + string(block.Blocks[count].Timestamp))
		fmt.Println("Purchase: " + block.Blocks[count].Purchase)
		fmt.Println("Cash: " + strconv.Itoa(block.Blocks[count].Cash))
		fmt.Println("Hash: " + block.Blocks[count].Hash)
		fmt.Println("Prev. hash: " + block.Blocks[count].PrevHash)
		fmt.Println("DIFF: " + strconv.Itoa(block.Blocks[count].Diff))
		fmt.Println("")
	}
}

func Add() {
	cash, err := strconv.ParseInt(os.Args[3], 10, 0)
	if err != nil {
		fmt.Println("Cash input parameter must be integer.\nUse for help: go run main.go help")
		os.Exit(1)
	}

	WriteFile(Filename, os.Args[2], int(cash))
}

func Mine() {
	diff, err := strconv.ParseInt(os.Args[2], 10, 0)
	if err != nil {
		fmt.Println("Mine input parameter must be integer.\nUse for help: go run main.go help")
		os.Exit(1)
	}
	Difficulty = int(diff)
}
