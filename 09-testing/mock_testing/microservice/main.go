package main

import (
	"fmt"
	"github.com/suhas1294/learn-go/09-testing/mock_testing/microservice/services"
)

func main() {
	allUsers := services.UserService.GetAllUsers()
	fmt.Println(allUsers)
}