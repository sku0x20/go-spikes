package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	clientRootCA := x509.NewCertPool()
	file, err := os.ReadFile("/home/exp/go-projects/go-spikes/tls/intermediate/certs/intermediate.cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	clientRootCA.AppendCertsFromPEM(file)
	cert, err := tls.LoadX509KeyPair(
		"/home/exp/go-projects/go-spikes/tls/intermediate/certs/san.example.com.fullchain-cert.pem",
		"/home/exp/go-projects/go-spikes/tls/intermediate/private/san.example.com.key.pem",
	)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    clientRootCA,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
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
		_, err = conn.Write([]byte("sending some data"))
		_, err = conn.Write([]byte(fmt.Sprintf("client name is %s", getClientName(conn))))
		if err != nil {
			log.Fatalf("write failed, %v", err)
		}
	}

}

func getClientName(conn net.Conn) string {
	conn2, ok := conn.(*tls.Conn)
	if !ok {
		log.Fatal("unable to downcast net.Conn to *tls.Conn")
	}
	emailId := conn2.ConnectionState().PeerCertificates[0].EmailAddresses[0]
	return emailId
}
