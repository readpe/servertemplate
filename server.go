package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// server type provides dependencies to handler methods
type server struct {
	router *http.ServeMux
}

// server constructor, should not setup many dependencies here
func newServer() *server {
	s := &server{router: http.NewServeMux()}
	s.routes() // assigns routes defined in routes.go to router
	return s
}

// ServeHTTP satisfies Handler interface so server can be used as a handler
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Handlers as methods

// Instead of a HandlerFunc, return HandlerFunc from wrapper func it is more flexible
func (s *server) handleIndex() http.HandlerFunc {
	// can place handler specific local variables here, also can pass in as parameters
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world! %s", r.URL.Path)
	}
}

func (s *server) handleAbout() http.HandlerFunc {
	// can place handler specific local variables here, also can pass in as parameters
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About me")
	}
}

// respond helper bare bones
func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.NotFound(w, r)
			return
		}
	}
}

// decode helper
func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// adminOnly is a middleware function which can wrapp a HandlerFunc
func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// example
		// if !currentUser(r).IsAdmin {
		// 	http.NotFound(w, r)
		// 	return
		// }
		h(w, r)
	}
}
