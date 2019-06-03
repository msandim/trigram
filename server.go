package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Server is ....
type Server struct {
	store *TrigramStore
}

// NewServer does ...
func NewServer(store *TrigramStore) *Server {
	return &Server{store: store}
}

func (server *Server) learnHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "/learn only accepts POST requests", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-type") != "text/plain" {
		http.Error(w, "/learn only accepts POST requests with 'text/plain' as 'Content-type'", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "invalid body received'", http.StatusMethodNotAllowed)
		return
	}

	body := string(bodyBytes)
	fmt.Println("GOT A LEARNING REQUEST: ", body)

	// Parse trigram:
	trigram := Trigram{"stuff1", "mariah", "brooklyn"}
	server.store.AddTrigram(trigram)
}

func (server *Server) generateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "/generate only accepts GET requests", http.StatusMethodNotAllowed)
		return
	}

	text := server.store.MakeText()

	fmt.Println("GENERATE THIS TEXT ", text)
	w.Write([]byte(text))
}

// Run does
func (server *Server) Run() {
	http.HandleFunc("/learn", server.learnHandler)
	http.HandleFunc("/generate", server.generateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
