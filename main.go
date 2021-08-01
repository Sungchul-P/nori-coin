package main

import (
	"github.com/Sungchul-P/nori-coin/cli"
	"github.com/Sungchul-P/nori-coin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
