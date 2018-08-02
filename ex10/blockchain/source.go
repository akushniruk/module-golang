package blockchain

const Filename = "data.json"

type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Purchase  string `json:"purchase"`
	Cash      int    `json:"cash"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevhash"`
}

type Blockchain struct {
	Blocks []Block `json:"blocks"`
}
