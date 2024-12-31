package storage

import (
	"io"
	"net/http"
)

type Server struct {
	AccessKey   string
	SecretKey   string
	EndpointURL string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewClient(s.AccessKey, s.SecretKey, s.EndpointURL)
	path := r.URL.Path[1:]
	switch r.Method {
	case http.MethodGet:
		b, err := c.Get(r.Context(), path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(b)
	case http.MethodPut:
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = c.Put(r.Context(), path, b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
