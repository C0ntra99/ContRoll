package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	serverMode := flag.Bool("server", false, "Starts ContRoll in server mode")
	port := flag.Int("port", 0, "The TCP port to connect on")

	flag.Parse()

	if *serverMode {
		contRollServer()
	}

	if *port < 1 || *port > 65535 {
		fmt.Print("Please specify a port to connect on with the --port=# flag")
		os.Exit(1)
	}

	fmt.Print("Connecting on port ", *port, "...")
	fmt.Print("This code doesn't actually do anything\nExiting...\n")
}

func contRollServer() {
	fmt.Print("ContRoll started in server mode...\n")
}
