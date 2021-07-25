package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

type PersonCreatedResponse struct{
	Name      string    `json:"name"`
	Job       string    `json:"job"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
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
	
	//Read the response body to cast it into model
	var createResponse PersonCreatedResponse
	if err = json.NewDecoder(resp.Body).Decode(&createResponse) ; err != nil{
		fmt.Println(err)
	}
	fmt.Println("Final response:\n\n", createResponse)
}
