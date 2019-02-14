package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/web/index.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe("localhost:80", nil)
}