package main

import (
	"bufio"
	"flag"
	"net"
	"fmt"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
	//fmt.Println("connection may be dead")
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
	for{
		connection,_:=ln.Accept()

		conns <- connection
		//fmt.Println("Connection made with",connection)
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
	reader:=bufio.NewReader(client)
	for{
		newMessage:=Message{}
		messageString,err:=reader.ReadString('\n')
		if err==nil{
			handleError(err)
		}

		newMessage.message=messageString
		newMessage.sender=clientid
		msgs<-newMessage
		//fmt.Println("Message recieved from",newMessage.sender,newMessage.message)
	}


}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.

	ln,_:=net.Listen("tcp",*portPtr)

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)


	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			var clientId int
			if len(clients)==0{
				clientId=0
			} else{
				clientId=len(clients)
			}
			clients[clientId]=conn
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client
			go handleClient(clients[clientId],clientId,msgs)

		case msg := <-msgs:
			for id,clientConnection:= range clients{
				if id!=msg.sender{
					//fmt.Println("message sent")
					fmt.Fprintln(clientConnection,msg.message)
				}

		}
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
		}
	}
}
