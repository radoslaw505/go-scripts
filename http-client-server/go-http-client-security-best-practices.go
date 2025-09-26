package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
)

func main() {
	certPool := x509.NewCertPool()
	pemData, _ := os.ReadFile("ca-cert.pem")
	certPool.AppendCertsFromPEM(pemData)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:    certPool,
			MinVersion: tls.VersionTLS12,
		},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://example.com")
	if err != nil {
		fmt.Println("TLS error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Secure TLS status:", resp.Status)
}
