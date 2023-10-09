package assignment01bca

import "fmt"

// Blockchain represents a simple blockchain.
type Blockchain struct {
	Blocks []*Block
}

// NewBlockchain creates a new blockchain with a genesis block.
func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis Block", 0, "")
	return &Blockchain{Blocks: []*Block{genesisBlock}}
}

// AddBlock adds a new block to the blockchain.
func (bc *Blockchain) AddBlock(transaction string, nonce int) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transaction, nonce, previousBlock.CurrentHash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// PrintBlocks prints all the blocks in the blockchain.
func (bc *Blockchain) PrintBlocks() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block #%d\n", i+1)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n\n", block.CurrentHash)
	}
}
