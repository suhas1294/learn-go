// https://stackoverflow.com/questions/34283255/closing-channel-of-unknown-length
package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"sync"
	"sort"
	"time"
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

type CommentTitle struct{
	NumberOfComments int `json:"number_of_comments"`
	Title string `json:"title"`
	Page int `json:"from_page"`
}

const (
	GET_HOST_URL = "https://jsonmock.hackerrank.com/"
	PATH = "api/articles"
)

var wg sync.WaitGroup

func main(){
	start := time.Now()
	comment_title_chan := make(chan CommentTitle)
	var commentTitleSlice []CommentTitle
	
	// pilot call to get total number of pages
	totalPage := makePaginatedRequest(1, comment_title_chan, true)

	// making concurrent requests to multiple pages at once
	for j:=1;j<=totalPage;j++{
		wg.Add(1)
		go makePaginatedRequest(j, comment_title_chan, false)
	}
	go func(){
		wg.Wait()
		close(comment_title_chan)
	}()
	for x:=range comment_title_chan {
		commentTitleSlice = append(commentTitleSlice,x)
	}
	sort.Sort(SortByNumberOfComments(commentTitleSlice))
	for _,article := range commentTitleSlice{
		fmt.Println(article.NumberOfComments, "\t\t", article.Title)
	}
	fmt.Println("\nTotal Number of valid articles:\t", len(commentTitleSlice))
	fmt.Println("Time taken:", time.Since(start))

	jd, _ := json.MarshalIndent(commentTitleSlice, "", "  ")
	fmt.Println(string(jd))

}

func makePaginatedRequest(pageNo int, chunk chan CommentTitle, pilotMode bool) int{
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
		defer wg.Done()
		for _, article := range articleResponse.Data{
			if(article.Title != "" && article.NumComments != 0){
				ct := CommentTitle{article.NumComments, article.Title, pageNo}
				chunk <- ct
			}
		}
	}
	return articleResponse.TotalPages
}

// sorting:
type SortByNumberOfComments []CommentTitle
func (c SortByNumberOfComments) Len() int {
	return len(c)
}
func (c SortByNumberOfComments) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c SortByNumberOfComments) Less(i, j int) bool {
	return c[i].NumberOfComments > c[j].NumberOfComments
}