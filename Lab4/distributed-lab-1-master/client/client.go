package main

import (
	"bufio"
	"flag"
	"net"
	"fmt"
	"os"
)

func read(conn *net.Conn) {
	for{
		reader:=bufio.NewReader(*conn)
		fmt.Println("message recieved from",conn)
		msg,_:=reader.ReadString('\n')
		fmt.Printf(msg)
	}
}

func write(conn *net.Conn) {
	fmt.Println("enter text to send:")
	for{
		stdin:=bufio.NewReader(os.Stdin)
		fmt.Println("enter text to send:")
		text,_:=stdin.ReadString('\n')
		fmt.Fprintln(*conn,text)

	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	conn,_:=net.Dial("tcp",*addrPtr)

	go write(&conn)
	go read(&conn)
	for{

	}
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
}
