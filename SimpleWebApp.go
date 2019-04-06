package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello Sai, how are you?")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "I'm good, thanks for asking")
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.ListenAndServe(":8443", nil)

}
