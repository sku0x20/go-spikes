package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
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
	revocationFile, err := os.ReadFile("/home/exp/go-projects/go-spikes/tls/intermediate/crl/intermediate.crl.der")
	revocationList, err := x509.ParseRevocationList(revocationFile)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    clientRootCA,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			clientCertSerialNumber := verifiedChains[0][0].SerialNumber
			for _, revokedCert := range revocationList.RevokedCertificates {
				if revokedCert.SerialNumber.Cmp(clientCertSerialNumber) == 0 {
					return errors.New("certificate is revoked")
				}
			}
			return nil
		},
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
			// this will fail when the client certificate is revoked
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
