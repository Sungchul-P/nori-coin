package main

import (
	"github.com/Sungchul-P/nori-coin/explorer"
	"github.com/Sungchul-P/nori-coin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
