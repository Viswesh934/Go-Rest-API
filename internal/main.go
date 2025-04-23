package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Idiot, %q", html.EscapeString(r.URL.Path))

	})
	log.Println("server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
