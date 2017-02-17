package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var url = "http://localhost/info"
	for {
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(string(data))
		time.Sleep(2 * time.Second)
	}

}
