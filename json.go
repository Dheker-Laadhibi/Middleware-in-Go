package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)
type fruitBasket struct{
Name string
Fruit []string
// `` altgr+7
ID int64 `json:"Ref"` // field appear as a key Ref 
/* 
In the Go programming language (Golang), 
fields that start with a lowercase letter in a struct are considered unexported or private. 
Unexported fields are not visible outside of the package they are defined in, */
private string // An unexported field  is not encoded 
Created time.Time


}
func main(){
	fmt.Println("hello json :::::::::::::::::::::::::::")
	
	basket:=fruitBasket{
		Name: "fff",
		Fruit: []string{"apple","banana","orange"},
		ID: 999,
		private: "seconde-rate",
		Created: time.Now(),
	}
	b,err :=json.Marshal(basket)
	if err !=nil{
		log.Println(err)
	}
	fmt.Println(b)
	// show keys and values
	fmt.Println(string(b))
	fmt.Println("pretty print")
//MarshalIndent another way to show json with indentation
/* {
- <=prefix             "Name": "fff",
 -            "Fruit": [
 -    4spaces   "apple",
 -                "banana",
-               "orange"
    ],
    "Ref": 999,
    "Created": "2024-01-31T16:46:52.7446895+01:00"
} */
b1,err := json.MarshalIndent(basket,"-","    ")
fmt.Println(string(b1))
/////////////////========decoding

var decod fruitBasket
// Unmarshal the JSON data into the struct
/* The reason why decoding works with b and not b1 
is because b1 contains the pretty-printed JSON string 
with indentation and line prefixes, 
which is not a valid JSON format. */
err1 := json.Unmarshal(b, &decod)
if err1 != nil {
	fmt.Println("Error unmarshaling JSON:", err)
	return
}

// Access the decoded data
	// Access the decoded data
	fmt.Println("Decoded Name:", decod.Name)
	fmt.Println("Decoded ID:", decod.ID)
	fmt.Println("Decoded Created:", decod.Created.Format(time.RFC3339))
     fmt.Println(decod)

///////////////////// arbitary object  and arrays 
child := []byte(`{"Name":"Eve", "Age":6}`)
var x interface{}
json.Unmarshal(child,&x)
fmt.Println(x)
//map[Age:6 Name:Eve]
// verifying if data has the type of 
data , ok:=x.(map[string]interface{})
if !ok {
	fmt.Println("Invalid data type")
	return
}

// Iterate over the map using a for loop with range
fmt.Println("Iterating over the Map:")
for key, value := range data {
	fmt.Printf("%s: %v\n", key, value)
}
}