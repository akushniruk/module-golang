package main

import (
	b "./blockchain"
	"fmt"
	"strconv"
)

func main() {
	b.WriteFile(b.Filename, "BIT TO ME:", 1)
	block := b.ReadFile(b.Filename)
	for count := 0; count < len(block.Blocks); count++ {
		fmt.Println("Index: " + strconv.Itoa(block.Blocks[count].Index))
		fmt.Println("Time: " + string(block.Blocks[count].Timestamp))
		fmt.Println("Purchase: " + block.Blocks[count].Purchase)
		fmt.Println("Cash: " + strconv.Itoa(block.Blocks[count].Cash))
		fmt.Println("Hash: " + block.Blocks[count].Hash)
		fmt.Println("Prev. hash: " + block.Blocks[count].PrevHash)
		fmt.Println("")
	}
}
