/*
package main

import (
	"fmt"
	"time"
	"net/http"
)

// current time function returns the current time in a specific format
func now() string {
	return time.Now().Format(time.Stamp) + ""
}

// The http.Request struct represents an HTTP request received by a server.
// This struct contains various fields, including information about
// the request method, URL, headers, form data, and more.

// When we use a pointer to an http.Request, we're essentially working with a reference
// to the actual request object, rather than a copy of it.

// handler1 handles HTTP requests, prints messages before and after processing
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(now() + "before")
	fmt.Println("handler 1")
	fmt.Println("after" + now())
}

// handler2 handles HTTP requests, prints messages before and after processing
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(now() + "before")
	fmt.Println("handler 2")
	fmt.Println("after" + now())
}

// handler3 handles HTTP requests, prints a message
func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler 3")
}

// logger is a middleware function to create a wrapper around other HTTP handlers
func logger(f http.HandlerFunc) http.HandlerFunc {
	// return an anonymous function
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(now() + "before")
		// wait until f is executed
		defer fmt.Println("after" + now())
		f(w, r)
	}
}

func main() {
	fmt.Println("starting")
	// When building a web app, there is probably some shared functionality
	// that you want to run from many or even all (http requests).
	// router => middleware => Application Handler 
	// Application Handler renders the response to the client 

	// Handle requests to "/h1" using handler1
	http.HandleFunc("/h1", handler1)

	// Handle requests to "/h2" using handler2
	http.HandleFunc("/h2", handler2)

	// Call middleware logger for requests to "/h3" with handler3
	http.HandleFunc("/h3", logger(handler3))

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}
*/
