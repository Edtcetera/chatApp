package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"sync"
)

type ServerEntry struct {
	addr net.Addr //This server
	next net.Addr //This server's invitee
	prev net.Addr //This server's inviter
}

type Server struct {
	entries map[string] *ServerEntry
	lock sync.Mutex
}

type Add_server_args struct {
	addr net.Addr //address of itself to be added to swarm's list of server
}

func main() {
	args := os.Args // ./server [inviterIP]
	if len(args) > 2 || len(args) < 1 {
		panic("Wrong arguments")
	}

	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic("TCP server creation error")
	}

	fmt.Println("Welcome to chat server, running server on " + ln.Addr().String())
	go tcpListener()
	switch len(args) {
	case 1:
		fmt.Println("You are the Origin chat server")
	case 2:
		fmt.Println("You are attempting to join a network invited by: " + args[1])
	}
}

func tcpListener(){
	server := new(Server)
	rpc.Register(server)
	rpc.HandleHTTP()
	//do something here so that we can accept both http and tcp connections
	l,e := net.Listen("tcp",":0")
	if e!=nil{
		log.Fatal("listen error:", e)
		return
	}
	go http.Serve(l, nil)
}

func (server *Server) addServer (args *Add_server_args, reply *string) error {
	addr := args.addr
	var err error

	server.lock.Lock()
	serverEntry := ServerEntry{addr, nil, nil}
	server.lock.Unlock()
	*reply = "Success"
	return err
}