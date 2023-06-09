package main

import (
	"crypto/tls"
	"log"
)

func main() {
	cert, err := tls.LoadX509KeyPair(
		"/home/exp/go-projects/go-spikes/tls/intermediate/certs/san.example.com.fullchain-cert.pem",
		"/home/exp/go-projects/go-spikes/tls/intermediate/private/san.example.com.key.pem",
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
			log.Fatal(err)
		}
		buff := make([]byte, 1024)
		_, err = conn.Read(buff) // add deadline?
		if err != nil {
			log.Fatalf("Read failed, %v", err)
		}
		log.Println(string(buff))
	}

}

// openssl s_client -servername www.example.com -connect localhost:14777 -CAfile tls/root/ca/certs/ca.cert.pem
