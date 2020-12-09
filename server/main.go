package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Launching Server...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection Failed... Err: %s", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // Connection Failed on func main()
		}
		time.Sleep(1 * time.Second)
	}
}
