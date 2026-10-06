package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type server struct {
	store *store
}

func newServer() *server {
	return &server{store: newStore()}
}

func (s *server) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/shorten", s.handleShorten)
	mux.HandleFunc("/", s.handleRedirect)
	return mux
}

func (s *server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"app": "tinyurl", "status": "ok"})
}

func (s *server) handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.URL == "" {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	code := s.store.save(body.URL)
	writeJSON(w, http.StatusCreated, map[string]string{"code": code, "url": body.URL})
}

func (s *server) handleRedirect(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" {
		writeJSON(w, http.StatusOK, map[string]string{"app": "tinyurl"})
		return
	}
	url, ok := s.store.get(code)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
