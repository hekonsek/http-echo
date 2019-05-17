package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/osexit"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
	osexit.ExitOnError(err)
}

func main() {
	fmt.Printf("Started echo server on %s .\n", color.GreenString("http://localhost:8080"))
	http.HandleFunc("/", handler)
	osexit.ExitOnError(http.ListenAndServe(":8080", nil))
}
