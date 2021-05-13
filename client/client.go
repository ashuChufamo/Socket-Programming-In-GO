package main


import (
	"bufio"
	"fmt"
	"net"
	"os"
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
		writter.WriteString(clientMsg+"\n")
		writter.Flush()

	}
}


func main()  {
	sock, err := net.Dial("tcp","127.0.0.1:8000")
	if err != nil{
		fmt.Println("an error occured, Couldn't connect to server.")
		fmt.Print(err)
		return
	}
	
	var wg sync.WaitGroup
	wg.Add(2)

	go incomingHandler(sock)
	go outgoingHandler(sock)
	wg.Wait()
	fmt.Println("in client, got here ...")

}