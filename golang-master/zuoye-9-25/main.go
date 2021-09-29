package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/zhang", k8s)

	log.Println("Starting first httpserver")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("golang++docker+kubernetes+istio"))
}
func k8s(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("kubernetes"))
}
