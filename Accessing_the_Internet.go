package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://www.w3schools.com/")
	bytes, _ := ioutil.ReadAll(resp.Body)

	string_Body := string(bytes)
	fmt.Println(string_Body)
	resp.Body.Close()
}
