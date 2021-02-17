package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(req http.ResponseWriter, res *http.Request) {
	req.Header().Set("Content-Type", "text/html")

	fmt.Fprint(req, "To get in touch, please send an email to <a href=\"mailto:test@mail.com\">test@mail.com</a>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":5000", nil)
}
