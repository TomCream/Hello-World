package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (stru Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, stru)
}

func (stri String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(stri))
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

