package main

import (
	"net"
	"os"
)

// https://www.golinuxcloud.com/golang-tcp-server-client/
const (
	ClientHost = "localhost"
	ClientPort = "8080"
	ClientType = "tcp"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr(ClientType, ClientHost+":"+ClientPort)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(ClientType, nil, tcpServer)
	if err != nil {
		println("Dial Failed: ", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte("This is a message"))
	if err != nil {
		println("Write date failed:", err.Error())
	}

	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read date failed:", err.Error())
		os.Exit(1)
	}
	println("Received message", string(received))

	conn.Close()
}
