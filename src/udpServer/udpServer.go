package udpServer

import (
	"fmt"
	"net"
)

const (
	host = "192.168.1.219"
)
func Start(port string) {
	addr, err := net.ResolveUDPAddr("udp",host+":"+port)
	if err != nil {
		fmt.Println("[!]Error could not resolve IP address!")
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("[!]Error could not start server!")
	}
	defer conn.Close()

	fmt.Printf("[+]Udp server started at %s\n", addr)
	bud := make([]byte, 1024)

	for {
		recv, addr, err := conn.ReadFromUDP(bud)
		fmt.Printf("%s -> %s\n", addr, string(bud[0:recv]))

		if err != nil {
			fmt.Println("[!]Error receiving!")
		}
	}

	//Write to client
	//conn.WriteTo([]byte("Hello from client"), conn.Network())

}
