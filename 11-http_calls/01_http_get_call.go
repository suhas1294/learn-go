package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const (
	GET_URL = "https://jsonmock.hackerrank.com/api/articles?page=1"
)

func main(){
	// step 1 : make a call
	resp, err := http.Get(GET_URL)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(resp)
	/*
		&{200 OK 200 HTTP/1.1 1 1 map[Access-Control-Allow-Origin:[*] 
		Content-Length:[2491] Content-Type:[application/json; 
		charset=utf-8] Date:[Fri, 02 Apr 2021 05:22:11 GMT] 
		Etag:[W/"9bb-vXhRlPnm1bo/tATKG0UHwQ"] X-Powered-By:[Express]] 
		0xc00009e040 2491 [] false false map[] 0xc000174000 0xc000379810}
	*/
	// fmt.Println(string(resp.Body)) // cannot convert resp.Body (type io.ReadCloser) to type string
	defer resp.Body.Close()

	// step 2 : converting 
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	sb := string(body)
	fmt.Printf(sb)
}

/* 
	func (c *Client) Get(url string) (resp *Response, err error)
	type Response struct{
		.
		.
		Body io.ReadCloser
		.
	}

	type ReadCloser interface {
		Reader
		Closer
	}

	type Reader interface {
		Read(p []byte) (n int, err error)
	}

*/