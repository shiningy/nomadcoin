package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/shiningy/nomadcoin/explorer"
	"github.com/shiningy/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the follwing flags:\n\n")
	fmt.Printf("-port: 		Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Sets the port of the server (default 4000)")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

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
