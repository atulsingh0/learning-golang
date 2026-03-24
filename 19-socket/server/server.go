package main

import (
	"log"
	"net"
)

func main() {

	conn, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Listening on :9999")
	}

	defer conn.Close()

	accept, err := conn.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer accept.Close()

	buf := make([]byte, 1024)
	n, err := accept.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	out := "Received at Server Side: " + string(buf[:n])
	log.Println(out)

	// return the data to conn
	n, err = accept.Write([]byte(out))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n, "bytes written successfully")
}
