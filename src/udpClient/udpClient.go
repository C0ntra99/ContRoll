package udpClient

import (
	"fmt"
	"net"
	//"time"
	//"strconv"
	"os"
	"bufio"
)

func Connect(host string, port string) {
	servAddr, err := net.ResolveUDPAddr("udp",host+":"+port)
	if err != nil {
		fmt.Println("[!]Error..")
	}
	//localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	//if err != nil {
	//	fmt.Println("[!]Error...")
	//}

	conn, err := net.DialUDP("udp", nil, servAddr)
	if err != nil {
		fmt.Println("[!]Error connecting!")
	}

	defer conn.Close()


	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		line := reader.Text()
		if line == "exit" {
			os.Exit(0)
		}

		buf := []byte(line+"\n")
		_,err := conn.Write(buf)
		if err != nil {
			fmt.Println("[!]Error sending...")
		}
	}
}
