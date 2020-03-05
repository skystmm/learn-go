package main

import (
	"log"
	"net"
	"os"
	"strings"
)

func handle(conn net.Conn) {
	b := make([]byte, 1024)
	for {
		conn.Read(b)
		if cap(b) > 0 {
			content := string(b)
			content = content[:strings.Index(content, "\r")]
			log.Printf(content)
			if content == "bye" {
				break
			}
			conn.Write(b)
		}
	}
	defer conn.Close()
}

func main() {
	listener, error := net.Listen("tcp", "0.0.0.0:8088")
	if error != nil {
		log.Println(error)
		os.Exit(2)
	}
	for {
		conn, error := listener.Accept()
		if error != nil {
			log.Println(error)
			continue
		}
		go handle(conn)
	}
}
