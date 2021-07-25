package services

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/suhas1294/learn-go/09-testing/mock_testing/microservice/models"
)

const url = "https://jsonplaceholder.typicode.com/users"
var UserService helper

type helper interface{
	GetAllUsers() *[]models.User
}

type UserHelper struct{}

func init(){
	UserService = &UserHelper{}
}

func (u *UserHelper) GetAllUsers() *[]models.User {
	
	// initialise container which holds response
	var users []models.User

	// make a request
	fmt.Println("Making a http request from service")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// decode the response stream into a variable that was initialised.
	dec := json.NewDecoder(resp.Body)
	for dec.More() {
		err = dec.Decode(&users)
		if err != nil {
			panic(err)
		}
	}
	return &users
}