package blockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/shiningy/nomadcoin/db"
	"github.com/shiningy/nomadcoin/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b.Data))
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{data, "", prevHash, height}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return block
}
