package main

import (
	"github.com/shiningy/nomadcoin/blockchain"
	"github.com/shiningy/nomadcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
