package main

import (

	 "fmt"
	 // to now time 
      "time"
// simple server with go 
	"net/http"




)
//current time 
func now() string {
	return time.Now().Format(time.Stamp) + ""
}
/* It represents an HTTP request

received by a server. This struct contains
 various fields, including information about
  the request method, URL, headers, form data,
   and more. */




// When we use a pointer to an http.Request, we're essentially working with a reference to the actual request object, rather than a copy of it.
//handler 1 start
func handler1(w http.ResponseWriter , r*http.Request){
	fmt.Println(now()+ "before")
	fmt.Println("handler 1")
	fmt.Println("after" + now())
}



//handler 2
func handler2(w http.ResponseWriter , r*http.Request){
	fmt.Println(now()+ "before")
	fmt.Println("handler 2")
	fmt.Println("after" + now())
}



 func main(){
	fmt.Println("starting")
// when you are building a web app there is probably some shared functionality that ypu want to run from many or even all (http requests)
//router=>middleware=>Application Handler 
// Application Handler render response to the client 





 }