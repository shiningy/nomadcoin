package main

import (
	"github.com/shiningy/nomadcoin/explorer"
	"github.com/shiningy/nomadcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
