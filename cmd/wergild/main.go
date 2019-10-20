package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/brianbroderick/wergild/internal/mud"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	service := ":2222"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()

	mud.GetServer().Start()
	listenForConnections(listener)
}

func listenForConnections(listener *net.TCPListener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		mud.NewDescriptor(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
