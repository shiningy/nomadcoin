package main

import (
	"github.com/shiningy/nomadcoin/cli"
	"github.com/shiningy/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
