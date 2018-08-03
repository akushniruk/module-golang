package blockchain

const Filename = "data.json"

var Difficulty int = 1

type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Purchase  string `json:"purchase"`
	Cash      int    `json:"cash"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevhash"`
	Diff      int    `json:"diff"`
	Nonce     string `json:"nonce"`
}

type Blockchain struct {
	Blocks []Block `json:"blocks"`
}
