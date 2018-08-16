package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	resp, err := http.Get("http://www.baidu.com/")

	if err != nil {

		// handle error

	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s",body)
}
