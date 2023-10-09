package main

import (
	"fmt"

	"github.com/HadiaTanveer/assignment01bca"
)

func main() {
	blockchain := assignment01bca.NewBlockchain()

	blockchain.AddBlock("Alice to Bob", 123)
	blockchain.AddBlock("Bob to Carol", 456)
	blockchain.AddBlock("Carol to Dave", 789)

	blockchain.PrintBlocks()

	// Optionally, you can add more blocks and print the blockchain again.

	isValid := VerifyChain(blockchain.Blocks)
	if isValid {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}

// VerifyChain verifies the integrity of the blockchain.
func VerifyChain(blocks []*assignment01bca.Block) bool {
	for i := 1; i < len(blocks); i++ {
		currentBlock := blocks[i]
		previousBlock := blocks[i-1]

		if currentBlock.PreviousHash != previousBlock.CurrentHash {
			return false
		}
	}
	return true
}
