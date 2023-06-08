package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// https://www.golinuxcloud.com/golang-tcp-server-client/
const (
	ServerHost = "localhost"
	ServerPort = "8080"
	ServerType = "tcp"
)

func main() {
	listen, err := net.Listen(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	println("got message", string(buffer))

	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprint("Your message is: ", string(buffer[:]), "Received at time: ", time)
	conn.Write([]byte(responseStr))

	conn.Close()
}
