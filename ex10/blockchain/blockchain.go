package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

//const target = 2

type Block struct {
	Index     int
	Timestamp string
	Cash      int
	Hash      string
	PrevHash  string
	// Proof of Work Target    int
	// Proof of Work Nonce     string
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + strconv.Itoa(block.Cash) + block.PrevHash + block.Timestamp
	sha := sha256.New()
	sha.Write([]byte(record))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

/* Proof of Work
func isHashValid(hash string, target int) bool {
	prefix := strings.Repeat("0", target)
	return strings.HasPrefix(hash, prefix)
}*/

func generateBlock(oldBlock Block, cash int) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Cash = cash
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	/*Proof of Work
	newBlock.Target = target
	for count := 0; ; count++ {
		hex := fmt.Sprintf("%d", count)
		newBlock.Nonce = hex
		if !isBlockValid(newBlock, oldBlock) {
			if !isHashValid(calculateHash(newBlock), newBlock.Target) {
				//fmt.Println(calculateHash(newBlock), "Noooooooooooooo")
				time.Sleep(time.Second / 100)
				continue
			} else {
				fmt.Println(calculateHash(newBlock), "OK!")
				newBlock.Hash = calculateHash(newBlock)
				break
			}
		}
	}
	Proof of Work*/
	Blockchain = append(Blockchain, newBlock)
	return newBlock
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index ||
		oldBlock.Hash != newBlock.PrevHash ||
		calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func genesisBlock() Block {
	t := time.Now()
	block := Block{}
	block = Block{0, t.String(), 0, calculateHash(block), ""}
	Blockchain = append(Blockchain, block)
	return block
}

func main() {
	block := genesisBlock()
	bit := generateBlock(block, 50)
	generateBlock(bit, 40)
	for _, node := range Blockchain {
		fmt.Println("Node:", node)
	}
}
