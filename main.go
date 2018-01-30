package main

import (
	"flag" // for taking arguments
	"./src/tcpServer"
	"./src/udpServer"
	"./src/tcpClient"
	"./src/udpClient"
)

func main() {

	// reads arguments
	listen := flag.Bool("listen", false, "Listen/Server mode")
	port := flag.String("port", "0", "The TCP port to connect on")
	host := flag.String("host", "127.0.0.1", "IP of the host you want to connect to")
	udp := flag.Bool("UDP", false, "Is this a UDP connection?")
	flag.Parse()

	switch *udp {
	case false: // TCP MODE
		switch *listen {
		case false:
			tcpClient.Connect(*host, *port)
		case true:
			tcpServer.Start(*port)
		default:
			flag.Usage()
		}

	case true: // UDP MODE
		switch *listen {
		case false:
			udpClient.Connect()
		case true:
			udpServer.Start()
		default:
			flag.Usage()
		}

	default:
		flag.Usage()
	}
}
