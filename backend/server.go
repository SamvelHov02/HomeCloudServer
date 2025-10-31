package backend

import (
	"fmt"
	"log"
	"net"
	"strconv"

	httphelper "github.com/SamvelHov02/HomeCloudHTTP"
)

var t = httphelper.Tree{}

func Start() {
	t.Init("Vault")
	t.Build()

	e := httphelper.EndPoint{}
	e.Get("/api/get/tree", GetTree)
	e.Get("/api/get", GetResource)
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

		req := httphelper.ReadRequest(conn)
		fn := e.Action(req.Method, req.Resource)

		req.Resource = apiToActualPath(req.Resource)

		resp := fn(req)

		_, err = conn.Write(resp)
		if err != nil {
			log.Fatal(err)
		}
		conn.Close()
	}
}

// Function for different endpoints
func GetTree(req httphelper.Request) []byte {
	t.ComputeHash()
	data, err := t.JSON()

	if err != nil {
		log.Fatal(err)
	}

	respHeader := httphelper.Header{
		"Content-Type":   []string{"application/json"},
		"Content-Length": []string{strconv.Itoa(len(data))},
	}

	data = httphelper.WriteResponse(data, httphelper.Status{Code: 200}, respHeader)
	return data
}

func GetResource(req httphelper.Request) []byte {
	data, status, respHeader := httphelper.ReadGetMethod(req.Resource, req.Headers)
	if status.Code != 200 {
		return nil
	}

	data = httphelper.WriteResponse(data, status, respHeader)
	return data
}

// API endpoint and actual path are different
// e.g. /api/get/file.txt -> /Vault/file.txt
func apiToActualPath(apiPath string) string {
	return apiPath[len("/api/get"):]
}
