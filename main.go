package main

import (
	"fmt"
	"net"
	"os"
)

var (
	ServerInstance *Server
	// BcryptCost is set lower when testing to speed up tests. It determines how much CPU
	// is used when hashing passwords
	BcryptCost = 14
)

func main() {
	service := ":2222"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()

	GetServer().Start()
	listenForConnections(listener)
}

func listenForConnections(listener *net.TCPListener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		newDescriptor(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
