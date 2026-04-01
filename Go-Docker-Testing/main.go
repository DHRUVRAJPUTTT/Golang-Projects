package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Dhruv and Our Docker is running great , %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hi again , this is also running great!!")
	})
	log.Fatal(http.ListenAndServe(":8081", nil))

}
