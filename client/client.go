package main

import (
	"log"
	"net"
	"os"
)

func main() {
	IDResponse, URLVideo := os.Args[1], os.Args[2]

	conn, err := net.Dial("tcp", "localhost:9999")

	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte(IDResponse + " " + URLVideo))

	conn.Close()
}
