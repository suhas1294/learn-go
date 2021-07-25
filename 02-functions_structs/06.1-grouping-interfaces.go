// this is a very absurd example just for sake of understanding grouping of interfaces
package main

import "fmt"
import "encoding/json"

// ********* two seperate interfaces *********
type person interface{
	speak() string
	think()
}

type human interface{
	reply(ping string) string
	compute()
}

// ********* grouping of above two interface *********
type homoSapiens interface{
	person
	human
}

// ********* struct having an interface as field *********
type Man struct{
	Gender string
	Masculine bool
}

type Women struct{
	Homo homoSapiens
	Gender string
	SourceOfBeauty bool
	Age int16
}

// ********* implementing all methods  *********
func (h *Man) speak() string{
	return fmt.Sprintf("I am speaking")
}

func (h *Man) think() {
	fmt.Println("I am thinking")
}

func (h *Man) compute() {
	fmt.Println("I am computing")
}

func (h *Man) reply(s string) string{
	return fmt.Sprintf(s, "pong")
}

// constructing a default value of type man
func createDefaultMan() *Man {
	return &Man{
		Gender: "male",
		Masculine: true,
	}
}

func main(){
	
	m := createDefaultMan()
	w := Women{
		// since m i.e., man is a struct that implements all methods 
		// of 'homoSapiens' interface, I can pass 'm' to 'Homo' field 
		// which is of type interface.
		Homo: m,
		Gender: "female",
		SourceOfBeauty: true,
		Age: 18,
	}
	jsonResp, _ := json.MarshalIndent(w, "", "  ")
 	fmt.Println(string(jsonResp))

	// Below example works:
	/*
		cannot use Women literal (type Women) as type homoSapiens in field value
		Women does not implement homoSapiens (missing compute method),
		If you plan to implements all methods on women struct, 
		refer commented block below at end of the file
	*/
	w2 := Women{
		Homo: Women{
			Gender: "female",
			SourceOfBeauty: false,
			Age: 23,
		},
		Gender: "female",
		SourceOfBeauty: true,
		Age: 18,
	}
	jsonResp2, _ := json.MarshalIndent(w, "", "  ")
 	fmt.Println(string(jsonResp2))
}

/* ignore from here, this is for test purpose
when a strcut is of embed type(i.e., it contains another struct / interface as field), 
you can not have pointer receivers attached to embed struct.

ref:
https://stackoverflow.com/questions/30403642/why-cant-the-interface-be-implemented-with-pointer-receivers
https://stackoverflow.com/a/40824044/5221796 */

/* func (w *Women) compute(){
	fmt.Println("women computing")
}

func(w *Women) reply(ping string) string{
	return fmt.Sprintf(ping, "pong from women")
}

func(w *Women) speak() string{
	return fmt.Sprintf("women speaking here..")
}

func (w *Women) think() {
	fmt.Prinln("women thinking..")
} */