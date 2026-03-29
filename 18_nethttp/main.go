package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 1 вариант
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//func main() {
//	http.HandleFunc("/", helloHandler)
//	http.ListenAndServe(":8080", nil) // http://localhost:8080
//}

// 2 вариант
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Custom handler")
}

func dsfd() {
	handler := MyHandler{}
	http.ListenAndServe(":8080", handler)
}

// маршрутизаторы
func marshruy() {
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", helloHandler).Methods("GET")
	r.HandleFunc("/users", helloHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	handler := http.HandlerFunc(helloHandler)
	http.ListenAndServe(":8080", loggingMiddleware(handler))
}
