package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my home page\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {

	fmt.Println()
	fmt.Println("Web server is starting on localhost, port 8080...")
	fmt.Println("To exit press Control-C.")
	fmt.Println()
	fmt.Println("You can try following urls:")
	fmt.Println("1) http://localhost:8080")
	fmt.Println("2) http://localhost:8080/hello")
	fmt.Println()

	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handler2)
	http.ListenAndServe(":8080", nil)
}
