package main

import (
	"net/http"
)

func main() {
	go func() {
		http.ListenAndServe(":8081", &fooHandler{})
	}()

	go func() {
		http.ListenAndServe(":8082", &zooHandler{})
	}()

	//the last call is outside goroutine to avoid that program just exit
	http.ListenAndServe(":8080", &barHandler{})
}

type fooHandler struct {
}

func (m *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on 8081: foo "))
}

type barHandler struct {
}

func (m *barHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on 8080: bar "))
}

type zooHandler struct {
}

func (m *zooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on 8082: zoo "))
}
