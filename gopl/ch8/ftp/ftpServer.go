package main

import (
	"log"
	"net"
	"os"
)

func main() {
	listener, error := net.Listen("tcp", "0.0.0.0:8223")
	if error != nil {
		log.Fatal(error)
		os.Exit(-1)
	}
	for {
		conn, error := listener.Accept()
		if error != nil {
			log.Fatalln(error)
			continue
		}
		go ftpHandle(conn)
	}
}

func ftpHandle(conn net.Conn) {

}
