package backend

import (
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

		resp := httphelper.ProcessRequest(conn)
		_, err = conn.Write(resp)
		if err != nil {
			log.Fatal(err)
		}
		conn.Close()
	}
}
