package main

import (
	"fmt"
	"log"
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
	args := os.Args // ./server [port]
	if len(args) > 3 || len(args) < 2 {
		panic("Wrong arguments, port and/or inviter IP required")
	}
	publicIP := getPublicIP()

	fmt.Println("Welcome to chat server, running server on port " + args[1])
	fmt.Println("Your public IP is: " + publicIP.String())

	switch len(args) {
	case 2:
		fmt.Println("You are the Origin chat server")
	case 3:
		fmt.Println("You are attempting to join a network invited by: " + args[2])
		//get server list from inviter
		handshakeServer(args)
	}
	createServer(args)
}

func handshakeServer(args []string) {
	conn, err := net.Dial("tcp", args[2])
	if err != nil {
		panic("handshake error")
	}
	fmt.Println(conn.LocalAddr().String())
}

func createServer(args []string) {
	ln, err := net.Listen("tcp", ":" + args[1])
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println(conn)
}

// Get preferred outbound ip of this machine
func getPublicIP() net.Addr {
	fmt.Println("Getting outbound IP addr")
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr()

	return localAddr
}