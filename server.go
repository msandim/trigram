package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
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

	err = server.learnText(body)

	if err != nil {
		http.Error(w, fmt.Sprintf("error while learning this text: %s", err), http.StatusBadRequest)
	}
}

func (server *Server) learnText(text string) error {

	// Get trigrams:
	trigrams, err := parseTrigrams(text)

	if err != nil {
		return err
	}

	for _, trigram := range trigrams {
		server.store.AddTrigram(trigram)
	}

	return nil
}

func parseTrigrams(text string) ([]Trigram, error) {

	// Remove any special characters and make all characters lower-case:
	text = strings.ToLower(regexp.MustCompile(`\.|,|;|!|\?`).ReplaceAllString(text, ""))

	words := strings.Split(text, " ")

	if len(words) < 3 {
		return nil, errors.New("text to learn needs to have more than 3 words")
	}

	var trigrams []Trigram

	for i := 0; i < len(words)-2; i++ {
		trigram := Trigram{words[i], words[i+1], words[i+2]}
		trigrams = append(trigrams, trigram)
	}

	return trigrams, nil
}

// GenerateHandler is
func (server *Server) generateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "/generate only accepts GET requests", http.StatusMethodNotAllowed)
		return
	}

	text := server.makeText()

	w.Write([]byte(text))
}

func (server *Server) makeText() string {
	return server.store.MakeText()
}

// Run does
func (server *Server) Run() {
	http.HandleFunc("/learn", server.learnHandler)
	http.HandleFunc("/generate", server.generateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
