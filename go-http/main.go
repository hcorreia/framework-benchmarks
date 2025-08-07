package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ok")
	})
	mux.HandleFunc("GET /db/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(w, "Not implemented!")
	})
	mux.HandleFunc("GET /chaos/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(w, "Not implemented!")
	})
	mux.HandleFunc("GET /health/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ok")
	})

	fmt.Println("Running on port", Env.Addr)
	log.Fatal(http.ListenAndServe(Env.Addr, mux))
}
