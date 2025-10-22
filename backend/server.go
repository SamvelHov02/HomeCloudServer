package backend

import (
	"fmt"
	"log"
	"net"

	httphelper "github.com/SamvelHov02/HomeCloudHTTP"
)

var t = httphelper.Tree{}

func Start() {
	t.Init("Vault")
	t.Build()

	e := httphelper.EndPoint{}
	e.Get("/Tree", GetTree)
	e.Get("Default", GetResource)
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

		responseHeader := httphelper.Header{}

		// resp := httphelper.ProcessRequest(conn)
		req := httphelper.ReadRequest(conn)

		var resp []byte

		if req.Method == "Get" {
			if fn, ok := e.GetEndpoints[req.Resource]; ok {
				bodyByte := fn(req)
				resp = httphelper.WriteResponse(bodyByte, httphelper.HTTPStatus{Code: 200}, responseHeader)
			} else {
				fn := e.GetEndpoints["Default"]
				bodyByte := fn(req)
				resp = httphelper.WriteResponse(bodyByte, httphelper.HTTPStatus{Code: 200}, responseHeader)
			}
		} else if req.Method == "Post" {

		} else if req.Method == "Put" {

		} else if req.Method == "Delete" {

		}

		_, err = conn.Write(resp)
		if err != nil {
			log.Fatal(err)
		}
		conn.Close()
	}
}

// Function for different endpoints
func GetTree(req httphelper.HTTPRequest) []byte {
	t.ComputeHash()
	data, err := t.JSON()

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func GetResource(req httphelper.HTTPRequest) []byte {
	data, status, _ := httphelper.ReadGetMethod(req.Resource, req.Headers)
	if status.Code != 200 {
		return nil
	}
	return data
}
