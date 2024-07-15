package server

import (
	"fmt"
	"net/http"

	skvError "github.com/bhkfazano/simple-kv/pkg/error"
	"github.com/bhkfazano/simple-kv/pkg/store"
)

type Server[K string, V string] struct {
	simpleKV   *store.SimpleKV[string, string]
	listenAddr string
}

func NewServer[K string, V string](simpleKV *store.SimpleKV[string, string], listenAddr string) *Server[string, string] {
	return &Server[string, string]{
		simpleKV:   simpleKV,
		listenAddr: listenAddr,
	}
}

func (s *Server[string, V]) ListenAndServe() error {
	fmt.Printf("Starting SimpleKV server\n")
	fmt.Printf("Server running on port %s\n", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, s)
}

func (s *Server[string, V]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found, use / to interact with SimpleKV", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		s.handleGet(w, r)
	case http.MethodPost:
		s.handlePost(w, r)
	case http.MethodDelete:
		s.handleDelete(w, r)
	case http.MethodPut:
		s.handlePut(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server[string, V]) handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling GET request\n")

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key not found in query parameters", http.StatusBadRequest)
		return
	}

	value, err := s.simpleKV.Get(key)
	if err != nil {
		if _, ok := err.(*skvError.NotFoundError); ok {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintf(w, "Key: %v, Value: %v\n", key, value)
}

func (s *Server[string, V]) handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling POST request\n")

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	if key == "" || value == "" {
		http.Error(w, "Key or value not found in query parameters", http.StatusBadRequest)
		return
	}

	err := s.simpleKV.Put(key, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Key: %v, Value: %v\n", key, value)
}

func (s *Server[string, V]) handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling DELETE request\n")

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key not found in query parameters", http.StatusBadRequest)
		return
	}

	value, err := s.simpleKV.Delete(key)
	if err != nil {
		if _, ok := err.(*skvError.NotFoundError); ok {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintf(w, "Key: %v, Value: %v\n", key, value)
}

func (s *Server[string, V]) handlePut(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling PUT request\n")

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	if key == "" || value == "" {
		http.Error(w, "Key or value not found in query parameters", http.StatusBadRequest)
		return
	}

	err := s.simpleKV.Put(key, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Key: %v, Value: %v\n", key, value)
}
