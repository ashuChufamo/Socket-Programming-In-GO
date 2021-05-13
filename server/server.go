package main


import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func incomingHandler(socket net.conn) {
	for {
		serverMsg, _ := bufio.NewReader(socket).ReadString("\n")
		fmt.Println("Server: ",serverMsg)	 
	}
}

func outgoingHandler(socket net.conn) {
	for{
		reader :=  bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		clientMsg := "this is a messaage from the client\n"
		writter:= bufio.NewWriter(sock)
		writter.WriteString(clientMsg + "\n")
		writter.Flush()

	}
}

func main()  {
	l, _ := net.Listen("tcp","127.0.0.1:8000")
	sock, _ := l.Accept()
	
	var wg sync.WaitGroup
	wg.Add(2)

	go incomingHandler(sock)
	go outgoingHandler(sock)

	wg.Wait()
	// fmt.Println("in server, got here ...")
}