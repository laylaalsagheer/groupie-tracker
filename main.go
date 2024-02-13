package main

import (
	"fmt"
	"log"
	"net/http"	
)

func main() {


	http.HandleFunc("/", APIHandler)
	http.HandleFunc("/details/", DetailsHandler)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Printf("Starting server for testing HTTP on Port 8080 ...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
