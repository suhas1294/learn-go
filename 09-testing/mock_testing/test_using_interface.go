package main

import "fmt"

type human interface{
	Ping() string
}

type male struct{}
type female struct{}

func (m *male) Ping() string{
	return "pong from male"
}

func (f *female) Ping() string{
	return "pong from female"
}

var personType human

func init(){
	// this is equivalent of performing a real action in model / service
	personType = &male{}
}

func main(){
	fmt.Println(personType.Ping())
	fmt.Println("Changing person type...")
	changePersonType()
	fmt.Println(personType.Ping())
}

// this is equivalent of mocking in your test file
func changePersonType(){
	personType = &female{}
}
