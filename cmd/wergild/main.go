package main

import (
	"net"
	"os"

	"github.com/brianbroderick/logit"
	"github.com/brianbroderick/wergild/internal/mud"
)

// 	Set the env DGRAPH_HOST if it's anything other than "127.0.0.1:9080"

func main() {
	service := ":2222"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()

	listenForConnections(listener)
}

func listenForConnections(listener *net.TCPListener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		mud.ServerInstance.AddConnection(conn)
	}
}

func checkError(err error) {
	if err != nil {
		logit.Fatal("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
