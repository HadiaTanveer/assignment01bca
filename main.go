package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Block represents a block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

// CalculateHash calculates the hash of a given string.
func CalculateHash(stringToHash string) string {
	// Convert the input string to bytes
	data := []byte(stringToHash)

	// Calculate the hash using SHA-256
	hash := sha256.Sum256(data)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash[:])

	return hashString
}

// NewBlock creates a new block with the given transaction, nonce, and previous hash.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	// Calculate the hash of the block's data
	blockData := fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)
	blockHash := CalculateHash(blockData)

	// Create the block
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		CurrentHash:  blockHash, // Assign the calculated hash to CurrentHash
	}

	return block
}

// CalculateHash calculates the hash of a given string.
func (b *Block) CalculateHash() string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// DisplayBlocks prints all the blocks in the blockchain.
func DisplayBlocks(blocks []*Block) {
	for i, block := range blocks {
		fmt.Printf("Block #%d\n", i+1)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n\n", block.CurrentHash)
	}
}

// ChangeBlock changes the transaction of a given block.
func ChangeBlock(block *Block, newTransaction string) {
	block.Transaction = newTransaction
	block.CurrentHash = block.CalculateHash()
}

// VerifyChain verifies the integrity of the blockchain.
func VerifyChain(blocks []*Block) bool {
	for i := 1; i < len(blocks); i++ {
		currentBlock := blocks[i]
		previousBlock := blocks[i-1]

		if currentBlock.PreviousHash != previousBlock.CurrentHash {
			return false
		}
	}
	return true
}

func main() {
	// Create the genesis block (the first block in the blockchain)
	genesisBlock := NewBlock("Genesis Block", 0, "")

	// Create some additional blocks
	block1 := NewBlock("Alice to Bob", 123, genesisBlock.CurrentHash)
	block2 := NewBlock("Bob to Carol", 456, block1.CurrentHash)
	block3 := NewBlock("Carol to Dave", 789, block2.CurrentHash)

	// Create a blockchain by adding blocks to a slice
	blocks := []*Block{genesisBlock, block1, block2, block3}

	// Display all blocks in the blockchain
	DisplayBlocks(blocks)

	// Change the transaction of block2
	ChangeBlock(block2, "New transaction from Bob to Carol") // Call the ChangeBlock function here

	// Display the blockchain after changing the transaction
	fmt.Println("______Blockchain after changing the transaction of block2:_____")
	DisplayBlocks(blocks)

	// Verify the integrity of the blockchain
	isValid := VerifyChain(blocks)
	if isValid {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}
