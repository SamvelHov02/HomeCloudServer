package backend

import (
	"bufio"
	"fmt"
	"log"
	"net"

	httphelper "github.com/SamvelHov02/HomeCloudHTTP"
)

func Start() {
	// Listen on port :8080 for connection
	l, err := net.Listen("tcp", ":8080")
	fmt.Println("Listening on port 8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		// Wait for connection request
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		request := getMessage(conn)

		httphelper.ProcessRequest(conn, request)
		conn.Close()
	}
}

func getMessage(conn net.Conn) string {
	// defer conn.Close()

	reader := bufio.NewReader(conn)

	var request string

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", err)
			request = message
			break
		}
	}
	return request
}
