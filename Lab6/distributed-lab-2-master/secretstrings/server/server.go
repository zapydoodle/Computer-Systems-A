
package main

import (
	"flag"
	"math/rand"
	"net"
	"net/rpc"
	"secretstrings/stubs"
	"time"
)

/** Super-Secret `reversing a string' method we can't allow clients to see. **/
func ReverseString(s string, i int) string {
    time.Sleep(time.Duration(rand.Intn(i))* time.Second)
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

type SecretStringOperations struct{}

func (s *SecretStringOperations) Reverse(request stubs.Request, response *stubs.Response)(err error){
	response.Message=ReverseString(request.Message,10)
	return
}

func (s *SecretStringOperations) FastReverse(request stubs.Request, response *stubs.Response)(err error){
	response.Message=ReverseString(request.Message,2)
	return
}


func main(){
	pAddr:=flag.String("port","8030","Port to listen on")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	rpc.Register(&SecretStringOperations{})
	listener,_:=net.Listen("tcp",":"+*pAddr)
	defer listener.Close()
	rpc.Accept(listener)





}
