package main
import (
//	"net/rpc"
	"flag"
	"net/rpc"
	"secretstrings/stubs"

	//	"bufio"
//	"os"
//	"secretstrings/stubs"
	"fmt"
)

func main(){
	server := flag.String("server","127.0.0.1:8030","IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	client,_:=rpc.Dial("tcp",*server)
	defer client.Close()

	request:=stubs.Request{Message:"Hello"}
	response:= new(stubs.Response)
	client.Call(stubs.ReverseHandler,request,response)
	fmt.Println("Responded: "+response.Message)




	//TODO: connect to the RPC server and send the request(s)
}
