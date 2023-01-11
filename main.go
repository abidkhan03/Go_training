package main

import (
	"fmt"
	"net/http"
	
)

func main() {
	http.HandleFunc("/hello", HelloHandler)
	fmt.Println("running server ")
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}
