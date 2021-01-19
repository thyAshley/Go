package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1>Welcome to Golang... Dynamic reloading</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":5000", nil)
}
