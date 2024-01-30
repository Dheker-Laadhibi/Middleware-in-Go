package main
import  g "github.com/gorilla/handlers"

import  "github.com/gorilla/mux"
import "fmt"
import "net/http"
import  "os"

func  homeHandler( w http.ResponseWriter , r *http.Request){
fmt.Println(w , "homepage ")
}
func  indexhandler( w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w , "indexpage ")
	}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/" , indexhandler)
	router.HandleFunc("/home" , homeHandler)
	//loginHandler  middleware of gorilla handlers 
	loggedRouter:= g.LoggingHandler(os.Stdout , router)
	http.ListenAndServe(":8080" , loggedRouter)
}
