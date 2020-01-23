// 3-3 스테이터스 코드 가져오기

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://localhost:18888")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	log.Println("status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)

	// Header도 출력 가능
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
	// 테스트 삼아 body의 길이를 출력해 봄
	log.Println("Length:", len(string(body)))
}
