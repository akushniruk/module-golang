package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + strconv.Itoa(block.Cash) + block.PrevHash + block.Purchase + block.Timestamp + block.Nonce
	sha := sha256.New()
	sha.Write([]byte(record))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

func isHashValid(hash string, target int) bool {
	prefix := strings.Repeat("0", target)
	return strings.HasPrefix(hash, prefix)
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index ||
		oldBlock.Hash != newBlock.PrevHash ||
		calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
func generateBlock(oldBlock Block, purchase string, cash int) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Purchase = purchase
	newBlock.Cash = cash
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	newBlock.Diff = diff
	for count := 0; ; count++ {
		hex := fmt.Sprintf("%d", count)
		newBlock.Nonce = hex
		if !isBlockValid(newBlock, oldBlock) {
			if !isHashValid(calculateHash(newBlock), newBlock.Diff) {
				time.Sleep(time.Second / 100)
				fmt.Println(calculateHash(newBlock), "NO")
				continue
			} else {
				fmt.Println(calculateHash(newBlock), "OK!")
				newBlock.Hash = calculateHash(newBlock)
				break
			}
		}
	}

	return newBlock
}

func genesisBlock() Block {
	t := time.Now()
	block := Block{}
	block = Block{0, t.String(), "", 0, calculateHash(block), "", diff, ""}
	return block
}

func (block *Blockchain) AddBlock(purchase string, cash int) {
	prevBlock := block.Blocks[len(block.Blocks)-1]
	newBlock := generateBlock(prevBlock, purchase, cash)
	block.Blocks = append(block.Blocks, newBlock)
}

func newBlockchain() *Blockchain {
	return &Blockchain{[]Block{genesisBlock()}}
}
