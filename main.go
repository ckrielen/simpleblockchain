package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"time"
)

// BlockChain All the blocks are collected in the Chain.
type BlockChain struct {
	Blocks []*Block
}

// AddBlock add a new block to the Chain
func (bc *BlockChain) AddBlock(b *Block) {
	bc.Blocks = append(bc.Blocks, b)
}

// CreateNewBlock Create a new block and add prev to chain
func (bc *BlockChain) CreateNewBlock(bPrev *Block) *Block {
	bPrev.hash = bPrev.createHash()
	b2 := &Block{Index: bPrev.Index + 1, Timestamp: time.Now(), prevHash: bPrev.hash}
	bc.AddBlock(bPrev)
	return b2
}

// Validate Validate the block chain
func (bc BlockChain) Validate() bool {
	prevHash := []byte{0}
	valide := true
	for _, b := range bc.Blocks {
		valide = valide && b.Validate(prevHash)
		prevHash = b.hash
		if !valide {
			break
		}
	}

	return valide
}

// ToString Create a String of the BlockChain usefull for debuging
func (bc BlockChain) ToString() string {
	result := "["
	for i, b := range bc.Blocks {
		result = fmt.Sprintf("%s %d (%s),", result, i, b.ToString())
	}

	return fmt.Sprintf("%s]", result)
}

// Block One Block of the block chain
type Block struct {
	Index     int       // index of the block
	Timestamp time.Time // Time stamp of the block
	BPM       int
	hash      []byte
	prevHash  []byte
}

// CreateHash Create the Hash for this block
func (b *Block) createHash() []byte {
	blockBodyHash := sha1.New()
	blockHash := sha1.New()
	// create a body with have the index, timestamp, BPM
	body := fmt.Sprintf("%d,%s,%d", b.Index, b.Timestamp, b.BPM)
	blockBodyHash.Write([]byte(body))
	blockHash.Write([]byte(append(b.prevHash, blockBodyHash.Sum(nil)...)))
	return blockHash.Sum(nil)
}

// Validate Check if the block is matching the hash
func (b Block) Validate(prevHash []byte) bool {
	if bytes.Compare(b.prevHash, prevHash) != 0 {
		return false
	}

	if bytes.Compare(b.hash, b.createHash()) != 0 {
		return false
	}
	return true
}

// ToString Create a String of the block usefull for debuging
func (b Block) ToString() string {
	return fmt.Sprintf("Index: %d, TimeStamp: %s, BPM: %d, Hash: %x, PrevHash: %x", b.Index, b.Timestamp, b.BPM, b.hash, b.prevHash)
}

func main() {
	blockChain := BlockChain{}
	firstBlock := &Block{Index: 0, Timestamp: time.Now(), BPM: 42, prevHash: []byte{0}}
	block2 := blockChain.CreateNewBlock(firstBlock)
	block2.BPM = 100
	block3 := blockChain.CreateNewBlock(block2)
	block3.BPM = 11
	fmt.Println("Block 1: ", firstBlock)
	fmt.Println("Block 2: ", block2)
	fmt.Println("Block 3: ", block3)
	fmt.Println("BlockChain", blockChain.ToString())
	fmt.Println("Blockchain is correct (expect true)", blockChain.Validate())
	block2.BPM = 200
	//	fmt.Println(blockChain.ToString())
	//fmt.Println(block2.ToString())
	fmt.Println("Blockchain is correct (expect false)", blockChain.Validate())
}
