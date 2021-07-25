// alternate library : https://github.com/yalp/jsonpath
package main

import(
	"fmt"
	"encoding/json"
	"github.com/oliveagle/jsonpath"
	"net/http"
	"io/ioutil"
)

func main(){
	const GET_URL = "https://jsonmock.hackerrank.com/api/articles?page=1"
	resp, err := http.Get(GET_URL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var json_data interface{}
	json.Unmarshal([]byte(respBody), &json_data)

	res, err := jsonpath.JsonPathLookup(json_data, "$.data[0].title")
	// res, err := jsonpath.JsonPathLookup(json_data, "$.data[?(@.num_comments > 0)].title")
	// res, err := jsonpath.JsonPathLookup(json_data, "$.data[?(@.story_url)].author") // story_url exists
	// res, err := jsonpath.JsonPathLookup(json_data, "$.data[0].num_comments")
	// res, err := jsonpath.JsonPathLookup(json_data, "$.data[0].created_at")
	// res, err := jsonpath.JsonPathLookup(json_data, "$.data")
	
	fmt.Printf("type is %T\t", res)
	fmt.Println(res)

}