package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello Omkar !<h1>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server Starting .... ")
	http.ListenAndServe(":2000", nil)
}
