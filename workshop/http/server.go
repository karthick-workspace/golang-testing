package main

import (
	"fmt"
	"log"
	"net/http"
)

func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi! I'm a new web server")
}

func RequestHandler() {
	http.HandleFunc("/", GetGreet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	RequestHandler()
}
