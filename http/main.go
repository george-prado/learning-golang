package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/", handler)
	fmt.Println("Listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hello, %s", name)
}
