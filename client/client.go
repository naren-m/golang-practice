// Reference https://tour.golang.org/methods/19
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ReqSample is sample request
type ReqSample struct {
	Name  string
	Title string
}

func main() {
	url := "http://localhost:8080"
	var requestBody ReqSample

	requestBody.Name = "Test name"
	requestBody.Title = "Samle title"

	// var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	jsonReq := []byte(fmt.Sprintf(`{"name":"%s", "title":"%s"`, requestBody.Name, requestBody.Title))
	fmt.Println(string(jsonReq))

	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonReq))

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("http.Do() error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() error: %v\n", err)
		return
	}

	fmt.Printf("read resp.Body successfully:\n%v\n", string(data))
}
