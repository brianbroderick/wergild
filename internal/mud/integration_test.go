package mud

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/brianbroderick/logit"
)

func init() {
	os.Setenv("PLATFORM_ENV", "test")
}

func TestConn(t *testing.T) {
	var wg sync.WaitGroup

	go server()

	wg.Add(1)
	go func() {
		var conn net.Conn
		var err error
		for {
			conn, err = net.Dial("tcp", ":2223")
			if err == nil {
				break
			}
		}

		// if _, err := c.Write([]byte("CONN AND LISTENER TEST")); err != nil {
		// 	t.Fatal(err)
		// }

		rb := make([]byte, 128)
		if _, err := conn.Read(rb); err != nil {
			t.Fatal(err)
		}

		str := string(rb)
		if strings.Contains(str, "What is your name:") {
			logit.Info("Hell ya")
		} else {
			logit.Info(fmt.Sprintf("|%s|", str))
		}

		wg.Done()
	}()

	wg.Wait()
}

func server() {
	service := ":2223"

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
		ServerInstance.AddConnection(conn)
	}
}

func checkError(err error) {
	if err != nil {
		logit.Fatal("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
