package mud

import (
	"net"
	"os"
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

		logit.Info(string(rb))

		// reader := bufio.NewReader(conn)
		// _, err = reader.ReadString('\n')
		// checkError(err)

		// for {
		// reader := bufio.NewReader(conn)
		// message, err := reader.ReadString('\n')
		// checkError(err)
		// logit.Info(message)
		logit.Info("HERE Me")
		// }
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
