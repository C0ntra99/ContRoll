package main

import (
	"flag" // for taking arguments
	"fmt"
	"os"
	"strings"
	"./src/tcpServer"
	"./src/udpServer"
	"./src/tcpClient"
	"./src/udpClient"
)

func main() {

	//All of this might want to be added to a seperate script....
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf(strings.Repeat("~", 50)+"\n")
		fmt.Printf("-u\t\t Sets the program in UDP mode\n")
		fmt.Printf("--listen\t\t Starts listening\n")
		fmt.Printf("--port={}\t\t Sets the port to Connect/Listen\n")
		fmt.Printf("--host={}\t\t Sets the host to connect to\n")
		}
	// reads arguments
	listen := flag.Bool("listen", false, "Listen/Server mode")
	port := flag.String("port", "0", "The TCP port to connect on")
	host := flag.String("host", "127.0.0.1", "IP of the host you want to connect to")
	udp := flag.Bool("u", false, "Is this a UDP connection?")

	flag.Parse()

	//Uses port to determine if no arguments have been passed
	if *port == "0" {
		flag.Usage()
		os.Exit(0)
	}

	//Logic is flawed, need to rewrite to that it will prin the usage...
	switch *udp {
	case false: // TCP MODE
		switch *listen {
		case false:
			tcpClient.Connect(*host, *port)
		case true:
			tcpServer.Start(*port)
		}

	case true: // UDP MODE
		switch *listen {
		case false:
			udpClient.Connect(*host, *port)
		case true:
			udpServer.Start(*port)
		}
	}
}
