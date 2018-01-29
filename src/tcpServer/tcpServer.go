package tcpServer

import (
  "fmt"
  "net"
  "os"
  "strconv"
  "bytes"
  "os/exec"
  "strings"
)

const (
  host = "192.168.1.219"
)

func StartServer(port int) {

  p := strconv.Itoa(port)
  listener, err := net.Listen("tcp",host+":"+p)
  if err != nil {
    fmt.Println("Error listening:", err.Error())
    os.Exit(1)
  }
  defer listener.Close()

  fmt.Println("Listening on " + host + ":" + p)
  for {
    conn, err := listener.Accept()
    if err != nil {
      fmt.Println("Error accecpting: ", err.Error())
    }

    fmt.Printf("Connection from %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())
    go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  buf := make([]byte, 1024)

  reqLen, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading: ", err.Error())
  }

  fmt.Println("Message len: ", strconv.Itoa(reqLen))
  msg := bytes.Index(buf, []byte{0})
  fmt.Println("Message: ", string(buf[:msg]))
  command := string(buf[:msg])
  ///If the command has a spcae in it then it will not run
  //Need to split the string somehow then send it to command
  //also need to do error checking
  cmd := exec.Command(command)
  output, _ := cmd.CombinedOutput()

  commandOutput(output,cmd)
  if len(output) > 0 {
    conn.Write([]byte(output))
  } else {
    conn.Write([]byte("Command Failed to execute"))
  }
  conn.Close()
}

func commandOutput(outs []byte, cmd *exec.Cmd) {
  fmt.Printf("Executing: %s\n", strings.Join(cmd.Args, " "))
  if len(outs) > 0 {
    //fmt.Printf("Output: %s\n", string(outs))
    fmt.Printf("%s Has been executed.", strings.Join(cmd.Args, " "))
  } else {
    fmt.Print("Command Failed to execute...\n")
  }
}
