package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func readReq(responseWriter http.ResponseWriter, request *http.Request) {

	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() error: %v\n", err)
		return
	}

	fmt.Printf("read request.Body successfully:\n%v\n", string(data))

	fmt.Fprintf(responseWriter, string("Hello Naren"))
}

func main() {
	http.HandleFunc("/", readReq)
	http.ListenAndServe(":8080", nil)
}
