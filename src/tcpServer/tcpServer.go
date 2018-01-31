package tcpServer

//Need to rewrite to be an actual server....
import (
  "fmt"
  "net"
  //"strconv"
  //"bytes"
  "bufio"
)

const (
  //Change to the localhost ip address
  host = "192.168.1.219"
)

func Start(port string) {

  listener, err := net.Listen("tcp",host+":"+port)
  if err != nil {
    fmt.Println("[!]Error couldnt listen on port " + port)
  }
  defer listener.Close()
  fmt.Printf("[+]Tcp server started at %s:%s\n", host, port)
  for {
    conn, err := listener.Accept()
    if err != nil {
      fmt.Println("[!]Error couldn't accept connection!")
    }
    go handleConnection(conn)
    }
  }

func handleConnection(conn net.Conn) {
  fmt.Printf("[+]Connection from %s", conn.RemoteAddr())
  conn.Write([]byte("~~~TCP server~~~\n"))

  //buffer := make([]byte, 1024)
  reader := bufio.NewReader(conn)

  for {
    bytes, err := reader.ReadString('\n')
    if err != nil {
      fmt.Printf("[!]Connection from %s has been closed\n", conn.RemoteAddr())
      return
    }
    fmt.Printf("%s", bytes)
  }

  //fmt.Print("[!]Connection from %s has been closed\n", conn.RemoteAddr())
  //conn.Read(buffer)





  //msg := bytes.Index(buffer, []byte{0})
  //fmt.Println("Message: ", string(buffer[:msg]))

}
