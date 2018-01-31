package udpClient

import (
	"fmt"
	"net"
	//"time"
	"strconv"
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

//This just sends the server numbers counting up
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		bud := []byte(msg)
		_,err := conn.Write(bud)
		if err != nil {
			fmt.Println(msg, err)
		}
		//time.Sleep(time.Second * 1)
	}
}
