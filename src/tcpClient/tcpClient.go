package tcpClient

import (
	"bufio"
	"fmt"
	"net"
	//"os"
)

func Connect(host string, port string) {
	fmt.Print("Connecting to ", host, " on port ", port, "...")
	address := host + ":" + port

	connection, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("[!]Error establishing the connection!")
	} else {
		fmt.Print("\nConnection established!\n")
	}

	reader := bufio.NewReader(connection)
	for {
		input, err := reader.ReadString('\n')
		if len(input)>0 {
			fmt.Println(input)
		}
		if err != nil {
			break
		}

		//fmt.Fprintf(connection, input+"\n")

		//response, _ := bufio.NewReader(connection).ReadString('\n')
		//fmt.Print(response)
	}
}
