package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Blockchain is our global blockchain.
var Blockchain []Block

// Block is our basic data structure!
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
}

// Main function
func main() {

}

// InitBlockchain creates our first Genesis node.
func InitBlockchain() {
	genesisBlock := Block{"Genesis Block", time.Now().Unix(), []byte{}, []byte{}}
	genesisBlock.Hash = genesisBlock.calculateHash()
	Blockchain = []Block{genesisBlock}
}

// NewBlock creates a new Blockchain Block.
func NewBlock(oldBlock Block, data string) Block {
	newBlock := Block{data, time.Now().Unix(), oldBlock.Hash, []byte{}}
	newBlock.Hash = newBlock.calculateHash()
	return newBlock
}

// AddBlock adds a new block to the Blockchain.
func AddBlock(b Block) error {
	// Validate block
	lastBlock := Blockchain[len(Blockchain)-1]
	if bytes.Compare(lastBlock.Hash, b.PrevHash) != 0 {
		return fmt.Errorf("Block is invalid. It should have PrevHash: %x, but has PrevHash: %x",
			lastBlock.Hash, b.PrevHash)
	}

	Blockchain = append(Blockchain, b)

	return nil
}

func (b *Block) calculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)
	headers := bytes.Join([][]byte{b.PrevHash, data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}
