package main

import (
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal("couldn's start server : ", err)
	}
	defer listener.Close()
	log.Print("Waiting to accept a connection")
	incomingConn, err := listener.Accept()
	if err != nil {
		log.Print("couldn't accept incoming conn: ", err)
		return
	}
	log.Print("Connection Accepted: ", incomingConn.RemoteAddr().String())
	var buf = make([]byte, 128)
	for {
		readBytes, err := incomingConn.Read(buf)
		if err != nil {
			log.Print("couldn't read properly due to: ", err)
			break
		}
		log.Print("Bytes Read: ", readBytes)
		_, err = incomingConn.Write(buf[:readBytes])
		if err != nil {
			log.Print("couldn't write properly due to :", err)
			break
		}
	}
	defer incomingConn.Close()
}
