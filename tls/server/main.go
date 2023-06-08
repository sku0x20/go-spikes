package main

import (
	"crypto/tls"
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
	listener, err := tls.Listen("tcp", "localhost:14777", cfg)
	if err != nil {
		log.Fatal(err)
	}
	_ = listener
}
