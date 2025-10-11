package main

import (
	"fmt"
	"server/backend"
)

func main() {
	fmt.Println("Main function from the server side")
	backend.Start()
}
