package cli

import (
	"flag"
	"fmt"
	"github.com/Sungchul-P/nori-coin/explorer"
	"github.com/Sungchul-P/nori-coin/rest"
	"os"
)

func usage() {
	fmt.Printf("Welcome to 노리 코인\n")
	fmt.Printf("Please use the following flags:\n")
	fmt.Printf("-port=4000:	set the PORT of the server\n")
	fmt.Printf("-mode=rest:	Choose between 'html' and 'rest'\n")
	os.Exit(0)
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
