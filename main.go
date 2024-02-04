//main.go

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	http.HandleFunc("/", APIHandler)
	fmt.Printf("Starting server for testing HTTP ...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
