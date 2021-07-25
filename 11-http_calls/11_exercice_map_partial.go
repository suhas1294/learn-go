package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"sync"
)

type ArticleResponse struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data []Article `json:"data"`
}

type Article struct {
	Title       string      `json:"title"`
	URL         string      `json:"url"`
	Author      string      `json:"author"`
	NumComments int         `json:"num_comments"`
	StoryID     int32 `json:"story_id"`
	StoryTitle  string `json:"story_title"`
	StoryURL    string `json:"story_url"`
	ParentID    int32 `json:"parent_id"`
	CreatedAt   int         `json:"created_at"`
}

type CommentTitleMap map[int]string

const (
	GET_HOST_URL = "https://jsonmock.hackerrank.com/"
	PATH = "api/articles"
)
var wait_group sync.WaitGroup

func main(){
	var commentTitleMapSlice []CommentTitleMap
	comment_title := make(chan map[int]string)
	
	// pilot call to get total number of pages
	totalPage := makePaginatedRequest(1, comment_title, true)
	wait_group.Add(totalPage)

	// making concurrent requests to multiple pages at once
	for j:=1;j<=totalPage;j++{
		go makePaginatedRequest(j, comment_title, false)
	}
	for ;totalPage >0; totalPage--{
		commentTitleMapSlice = append(commentTitleMapSlice, <- comment_title)
	}
	for _,val := range commentTitleMapSlice{
		for nc, title := range val{
			fmt.Println("Number Of comments:", nc, "\t\tTitle:\t", title)
		}
	}
}

func makePaginatedRequest(pageNo int, chunk chan map[int]string, pilotMode bool) int{
	fmt.Println("Making request to page: ", pageNo)
	uri := GET_HOST_URL + PATH
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println(err)
	}
	q := req.URL.Query()
	q.Add("page", strconv.Itoa(pageNo))
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        fmt.Println("Error on response.\n[ERROR] -", err)
    }
	defer resp.Body.Close()
	var articleResponse ArticleResponse
	if err = json.NewDecoder(resp.Body).Decode(&articleResponse) ; err != nil{
		fmt.Println(err)
	}
	if !pilotMode{
		m := make(map[int]string)
		for _, article := range articleResponse.Data{
			if(article.Title != "" && article.NumComments != 0){
				m[article.NumComments] = article.Title
			}
		}
		chunk <- m
		wait_group.Done()
	}
	fmt.Println("")
	return articleResponse.TotalPages
}