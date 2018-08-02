package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + strconv.Itoa(block.Cash) + block.PrevHash + block.Purchase + block.Timestamp
	sha := sha256.New()
	sha.Write([]byte(record))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
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
	block = Block{0, t.String(), "", 0, calculateHash(block), ""}
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
