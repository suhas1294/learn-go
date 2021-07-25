package main

import (
	"fmt"
	"encoding/json"
)

type User struct{
	Name string
	Age int
}

type AllUsers struct{
	Everyone []User
}

var all AllUsers

func main(){

	var all_data = `[
	{
		"Name": "someOne-1",
		"Age": 23
	},
	{
		"Name": "someOne-3",
		"Age": 89
	},
	{
		"Name": "someOne-2",
		"Age": 56
	}
]`


	err := json.Unmarshal( []byte(all_data), &all.Everyone )
	if (err != nil){
		fmt.Println("exiting")
	}

	//fmt.Println(all.Everyone)
	for _, x := range all.Everyone{
		fmt.Println(x.Name, ",\t", x.Age)
	}

	fmt.Println(len(all.Everyone))

}

