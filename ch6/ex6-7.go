// OS 에 의존하지 않은 채 인증서를 읽어와 HTTPS와 통산하는 클라이언트

package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	cert, err := ioutil.ReadFile("./ca.crt")

	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	tlsConfig := &tls.Config{
		//RootCAs: certPool,
		InsecureSkipVerify: true,
	}

	tlsConfig.BuildNameToCertificate()

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	resp, err := client.Get("https://localhost:18443")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}

	log.Println(string(dump))
}

/*

openssl genrsa -out ca.key 2048

openssl req -new -sha256 -key ca.key -out ca.csr -config openssl.cnf

openssl x509 -in ca.csr -days 365 -req -signkey ca.key -sha256 -out ca.crt -extfile ./openssl.cnf -extensions CA
---

openssl genrsa -out server.key 2048

openssl req -new -nodes -sha256 -key server.key -out server.csr -config openssl.cnf

openssl x509 -req -days 365 -in server.csr -sha256 -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions Server

*/
