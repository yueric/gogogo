package main

import (
	"net/http"
	"fmt"
	"log"
	"html"
)
func Indexhandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hello world")
}
func main(){
	http.HandleFunc("/",Indexhandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello, %q, %v", html.EscapeString(r.URL.Path),r.URL.Query())

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}