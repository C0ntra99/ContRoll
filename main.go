package main

import (
	"flag"
	"fmt"
)

func main() {

	serverMode := flag.Bool("server", false, "Starts ContRoll in server mode")
	port := flag.Int("port", 0, "The TCP port to connect on")
	host := flag.String("host", "127.0.0.1", "IP of the host you want to connect to")
	flag.Parse()

	if *port < 1 || *port > 65535 {
		fmt.Print("Please specify a port to connect on with the --port=# flag")
	}

	if *serverMode {
		contRollServer(*port)
	} else {
		contRollClient(*host, *port)
	}

	fmt.Print("\n\nThis code doesn't actually do anything\nExiting...\n")
}

func contRollServer(port int) {
	fmt.Print("ContRoll started in server mode...\n\n")
	fmt.Print("Waiting for connections on port ", port, "...")
}

func contRollClient(host string, port int) {
	fmt.Print("ContRoll started in client mode...\n\n")
	fmt.Print("Connecting to ", host, " on port ", port, "...")
}
