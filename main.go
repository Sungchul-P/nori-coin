package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcom to 노리 코인\n")
	fmt.Printf("Pleas use the following commands:\n")
	fmt.Printf("explorer:		Start the HTML Explorer\n")
	fmt.Printf("rest:			Start the REST API (recommended)\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		fmt.Println("Start REST API")
	default:
		usage()
	}

}
