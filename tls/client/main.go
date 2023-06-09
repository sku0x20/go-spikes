package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
)

func main() {
	roots := x509.NewCertPool()
	file, err := os.ReadFile("/home/exp/go-projects/go-spikes/tls/root/ca/certs/ca.cert.pem")
	roots.AppendCertsFromPEM(file)
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{
		RootCAs:    roots,
		ServerName: "san.example.com",
	}
	conn, err := tls.Dial("tcp", "localhost:14777", config)
	if err != nil {
		log.Fatalf("error dialing %v", err)
	}
	_, err = conn.Write([]byte("some data"))
	if err != nil {
		log.Fatalf("write failed, %v", err)
	}
	buff := make([]byte, 1024)
	_, err = conn.Read(buff) // add deadline?
	if err != nil {
		log.Fatalf("read failed, %v", err)
	}
	log.Printf("received data: %v\n", string(buff))

}
