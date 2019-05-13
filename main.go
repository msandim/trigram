package main

import (
	"fmt"
	"log"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

func getOnly(h handler) handler {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)
			return
		}
		http.Error(w, "get only", http.StatusMethodNotAllowed)
	}
}

func postOnly(h handler) handler {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}

func learn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I learn %s!", r.URL.Path[1:])
}

func generate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I generate %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/learn", postOnly(learn))
	http.HandleFunc("/generate", getOnly(generate))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
