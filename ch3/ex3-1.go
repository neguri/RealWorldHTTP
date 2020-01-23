// GET 메서즈로 액세스해서 바디를 획득

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18888")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	log.Println(string(body))

}
