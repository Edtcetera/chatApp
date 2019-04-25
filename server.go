package main

import (
	"fmt"
	"net"
	"os"
)

type server struct {
	addr net.Addr
	next net.Addr
	prev net.Addr
}

var serverList []server

func main() {
	args := os.Args // ./server [inviterIP]
	if len(args) > 2 || len(args) < 1 {
		panic("Wrong arguments")
	}

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		panic("TCP server creation error")
	}

	fmt.Println("Welcome to chat server, running server on " + ln.Addr().String())

	switch len(args) {
	case 1:
		fmt.Println("You are the Origin chat server")
	case 2:
		fmt.Println("You are attempting to join a network invited by: " + args[1])
		//get server list from inviter
		handshakeServer(args)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}


}

func handshakeServer(args []string) {
	conn, err := net.Dial("tcp", args[1])
	if err != nil {
		panic("handshake error")
	}
	fmt.Println(conn.LocalAddr().String())
}

func handleConnection(conn net.Conn) {
	fmt.Println(conn.LocalAddr().String())
	fmt.Println(conn.RemoteAddr().String())
}