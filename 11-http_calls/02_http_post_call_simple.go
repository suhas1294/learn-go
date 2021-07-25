package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

const POST_URL = "https://reqres.in/api/users"

func main() {
	reqPayload := Person{
		"Foobar",
		"Developer",
	}
	jsonPayloadBs, _ := json.Marshal(reqPayload)
	reqBufferReader := bytes.NewBuffer(jsonPayloadBs)
	resp, err := http.Post(POST_URL, "application/json", reqBufferReader)
	if err != nil {
		fmt.Println("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	sb := string(body)
	fmt.Println(sb)
}
