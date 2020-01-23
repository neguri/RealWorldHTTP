// 3-2 오류 체크를 생략한 더 작은 코드

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
}
