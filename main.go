package main

import (
	"fmt"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello Go"))
	fmt.Println("Hello For me")
}

func main() {
	fmt.Println("Webserver started")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
