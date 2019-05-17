package main

import (
	"net/http"

	_ "net/http/pprof"

	"time"
)

var Sleep = time.Sleep

func main() {
    http.ListenAndServe(addr, handler)
    http.ListenAndServe(a)
}
