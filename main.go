package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"./src/tcpServer"
)

func main() {

	serverMode := flag.Bool("server", false, "Starts ContRoll in server mode")
	port := flag.Int("port", 0, "The TCP port to connect on")
	host := flag.String("host", "127.0.0.1", "IP of the host you want to connect to")
	udp := flag.Bool("UDP", false, "Is this a UDP connection?")
	flag.Parse()

	if *port < 1 || *port > 65535 {
		fmt.Print("Please specify a port to connect on with the --port flag\n\n")
		os.Exit(1)
	}

	if *serverMode {
		contRollServer(*port)
	} else {
		contRollClient(*host, *port, *udp)
	}

	fmt.Print("\n\nThis code doesn't actually do anything\nExiting...\n")
}

func contRollServer(port int) {
	fmt.Print("ContRoll started in server mode...\n\n")
	tcpServer.StartServer(port)
}

func contRollClient(host string, port int, udp bool) {
	fmt.Print("ContRoll started in client mode...\n\n")
	fmt.Print("Connecting to ", host, " on port ", port, "...")

	address := host + fmt.Sprintf("%v", port)
	protocol := "tcp"
	if udp {
		protocol = "udp"
	}

	connection, _ := net.Dial(protocol, address)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nConnection established!\n")

		input, _ := reader.ReadString('\n')
		fmt.Fprintf(connection, input+"\n")

		response, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Print(response)
	}

}
