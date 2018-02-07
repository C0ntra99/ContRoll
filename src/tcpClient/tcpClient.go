package tcpClient

import (
	"bufio"
	"fmt"
	"net"
	"os"
)


func Connect(host string, port string) {
	fmt.Print("[-]Connecting to ", host, " on port ", port, "...")
	address := host + ":" + port

	connection, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("[!]Error establishing the connection!")
		os.Exit(0)
	} else {
		fmt.Print("\n[+]Connection established!\n")
	}

	servReader := bufio.NewReader(connection)
	//Wont move on until input from server...
	for {
		input, err := servReader.ReadString('\n')
		if len(input)>0 {
			fmt.Println(input)
		}
		if err != nil {
			break
		}

		//Send the server cool shit
		clientReader := bufio.NewScanner(os.Stdin)
		for clientReader.Scan() {
			line := clientReader.Text()
			if line == "exit" {
				os.Exit(0)
			}


			buf := []byte(line+"\n")
			_,err := connection.Write(buf)
			if err != nil {
				fmt.Println("[!]Error sending...")
			}
		}
	}
}
