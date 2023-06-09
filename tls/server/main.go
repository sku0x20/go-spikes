package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	cert, err := tls.LoadX509KeyPair(
		"/home/exp/go-projects/go-spikes/tls/intermediate/certs/www.example.com-chain.cert.pem",
		"/home/exp/go-projects/go-spikes/tls/intermediate/private/www.example.com.key.pem",
	)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":14777", cfg)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		_, err = conn.Write([]byte("some data"))
		if err != nil {
			return
		}
	}

}
