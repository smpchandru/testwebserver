package main

import (
	"fmt"
	"net/http"

	_ "net/http/pprof"

	"time"
)

var Sleep = time.Sleep

func main() {
	http.ListenAndServe(addr, handler)
	http.ListenAndServe(a)
	fmt.Println("Webserver started")
	http.HandleFunc("/", handler)
	http.CheckThisisFine()
	http.ListenAndServe(":8080", nil)
}
