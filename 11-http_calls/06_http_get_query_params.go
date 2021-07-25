package main

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"bytes"
	"net/http/httputil"
)

const (
	GET_URL = "https://jsonmock.hackerrank.com/api/articles"
)

type ArticleRequest struct{
	Type string
	Pages int16
}

func main(){

	// step-1 : preparing the request
	a := ArticleRequest{
		"article",
		1,
	}
	paylaodBs, _ := json.Marshal(a)
	reqBufferReader := bytes.NewBuffer(paylaodBs)
	req, err := http.NewRequest("GET", GET_URL, reqBufferReader)
	req.Header.Add("Authorization", "Bearer AbCdEf123456")

	q := req.URL.Query()
	// q.Add("page", 5) //cannot use 5 (type untyped int) as type string in argument to q.Add
	q.Add("page", "5")
	q.Add("orderBy", "desc")
	req.URL.RawQuery = q.Encode()

	// making the request
	client := &http.Client{}
	resp, err := client.Do(req)
	
	// interceptor.FormatRequest(req)
	requestDump, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(requestDump))

	if err != nil {
        fmt.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()

	// converting response to string and printing
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	sb := string(respBody)
	fmt.Printf(sb)

}