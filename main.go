package main

import (
	"fmt"
	"log"
	"net/http"
)
	
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello"
}


func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	htttp.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}