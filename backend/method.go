package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	httphelper "github.com/SamvelHov02/HomeCloudHTTP"
)

const VaultPath = "/home/samo/dev/HomeCloud/server"

// Function returns the requested file
func GetFile(req httphelper.Request) ([]byte, httphelper.Status, httphelper.Header) {
	resp := httphelper.Body{}
	respHeader := httphelper.Header{}
	var Status httphelper.Status
	fileData, err := os.ReadFile(VaultPath + req.Resource)

	if err != nil {
		log.Fatal(err)
	}

	Status.Code = 200

	for _, key := range req.Headers.Keys() {
		switch key {
		case "Accept":
			val, _ := req.Headers.Get(key)
			if val[0] == "application/json" {
				resp.Data = string(fileData)
			}
		}
	}

	dataBytes, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	respHeader.Add("Content-Type", "application/json")
	respHeader.Add("Content-Length", strconv.Itoa(len(dataBytes)))
	return dataBytes, Status, respHeader
}

func PostFile(req httphelper.Request) ([]byte, httphelper.Status, httphelper.Header) {
	var Status httphelper.Status
	var RespHeader httphelper.Header

outer:
	for _, key := range req.Headers.Keys() {
		switch key {
		case "Content-Type":
			// Only accept application/json, API server
			if h, _ := req.Headers.Get(key); h[0] != "application/json" {
				Status.Code = 400
				break outer
			}
		case "Content-Length":
			val, _ := req.Headers.Get(key)
			lenVal, err := strconv.Atoi(val[0])

			// Content can't be negative length or not numeric
			if lenVal <= 0 || err != nil {
				Status.Code = 400
				break outer
			}
		}
	}

	if Status.Code == 0 {
		if _, err := os.Stat(VaultPath + req.Resource); os.IsNotExist(err) {
			file, err := os.Create(VaultPath + req.Resource)

			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			file.WriteString(req.Data.Data)
			Status.Code = 204
		} else {
			Status.Code = 409
		}
	}

	RespHeader.Add("Content-Length", "0")
	RespHeader.Add("Content-Type", "application/json")
	return []byte{}, Status, RespHeader
}

func PostDir(req httphelper.Request) ([]byte, httphelper.Status, httphelper.Header) {
	var Status httphelper.Status
	var RespHeader httphelper.Header

outer:
	for _, key := range req.Headers.Keys() {
		switch key {
		case "Content-Type":
			// Only accept application/json, API server
			if h, _ := req.Headers.Get(key); h[0] != "application/json" {
				Status.Code = 400
				break outer
			}
		case "Content-Length":
			val, _ := req.Headers.Get(key)
			lenVal, err := strconv.Atoi(val[0])

			// Content can't be negative length or not numeric
			if lenVal <= 0 || err != nil {
				Status.Code = 400
				break outer
			}
		}
	}

	if Status.Code == 0 {
		fmt.Println(VaultPath + req.Resource)
		if _, err := os.Stat(VaultPath + req.Resource); os.IsNotExist(err) {
			err = os.Mkdir(VaultPath+req.Resource, 0755)

			if err != nil {
				log.Fatal(err)
			}

			Status.Code = 204
		} else {
			Status.Code = 409
		}
	}

	RespHeader.Add("Content-Length", "0")
	RespHeader.Add("Content-Type", "application/json")

	return []byte{}, Status, RespHeader
}

func PutFile(req httphelper.Request) ([]byte, httphelper.Status, httphelper.Header) {
	var respHeader httphelper.Header
	var Status httphelper.Status

outer:
	for _, key := range req.Headers.Keys() {
		switch key {
		case "Content-Type":
			if h, _ := req.Headers.Get(key); h[0] != "application/json" {
				Status.Code = 400
				break outer
			}
		case "Content-Length":
			val, _ := req.Headers.Get(key)
			lenVal, err := strconv.Atoi(val[0])

			if lenVal <= 0 || err != nil {
				Status.Code = 400
				break outer
			}
		}
	}

	if Status.Code == 0 {
		if _, err := os.Stat(VaultPath + req.Resource); err == nil {
			err = os.WriteFile(VaultPath+req.Resource, []byte(req.Data.Data), 0644)

			if err != nil {
				Status.Code = 400
			}

			Status.Code = 200
		}
	}

	respHeader.Add("Content-Type", "application/json")
	respHeader.Add("Content-Length", "0")

	return []byte{}, Status, respHeader
}
