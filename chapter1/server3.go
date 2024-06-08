package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"	
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for key, value := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", key, value)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for key, value := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", key, value)
	}
}