package tcpClient

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Connect(host string, port string) {
	fmt.Print("Connecting to ", host, " on port ", port, "...")
	address := host + "" + port

	connection, _ := net.Dial("tcp", address)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nConnection established!\n")

		input, _ := reader.ReadString('\n')
		fmt.Fprintf(connection, input+"\n")

		response, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Print(response)
	}
}
