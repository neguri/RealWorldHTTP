// OS 에 의존하지 않은 채 인증서를 읽어와 HTTPS와 통산하는 클라이언트

package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")

	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
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

openssl genrsa -out client.key 2048

openssl req -new -nodes -sha256 -key client.key -out client.csr -config openssl.cnf

openssl x509 -req -days 365 -in client.crt -sha256 -out client.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions Client
---
*/
