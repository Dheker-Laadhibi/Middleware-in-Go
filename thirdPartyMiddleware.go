package main

import (
	"fmt"
	"net/http"
	"os"

	g "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage ")
}
func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "indexpage ")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexhandler)
	router.HandleFunc("/home", homeHandler)
	//loginHandler  middleware of gorilla handlers
	loggedRouter := g.LoggingHandler(os.Stdout, router)
	http.ListenAndServe(":8080", loggedRouter)
}
