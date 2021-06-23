package main

import (
	"fmt"
	"github.com/Sungchul-P/nori-coin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}
