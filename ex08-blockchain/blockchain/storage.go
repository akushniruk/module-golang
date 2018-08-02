package blockchain

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadFile(fileName string) Blockchain {
	jsonFile, err := os.Open(fileName)
	checkErr(err)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var block Blockchain

	json.Unmarshal(byteValue, &block)

	return block
}

func createFile() {
	block1 := newBlockchain()

	jsonFile, err := os.Create(Filename)
	checkErr(err)

	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	err = encoder.Encode(&block1)
	checkErr(err)
}

func fileExist(filename string) bool {
	if _, err := os.Open(filename); err == nil {
		return false
	}
	return true
}

func WriteFile(filename string, purchase string, cash int) {
	block := ReadFile(filename)
	if fileExist(filename) || block.Blocks == nil {
		createFile()
	} else {
		jsonFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
		checkErr(err)
		jsonWriter := io.Writer(jsonFile)
		encoder := json.NewEncoder(jsonWriter)
		block.AddBlock(purchase, cash)
		err = encoder.Encode(&block)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Massange", err)
	}
}
