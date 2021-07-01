package main

import (
	"github.com/Sungchul-P/nori-coin/blockchain"
	"github.com/Sungchul-P/nori-coin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
