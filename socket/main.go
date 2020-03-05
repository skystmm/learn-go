package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func handle(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		b := input.Text()
		if len(b) > 0 {
			content := b
			log.Printf(content)
			if content == "bye" {
				break
			}
			conn.Write([]byte(b))

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
