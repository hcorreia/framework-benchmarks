package main

import (
	"fmt"
	"framework-benchmarks/go-http/utils"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, http.StatusOK, utils.JSON{
			"result": "Ok",
		})
	})
	mux.HandleFunc("GET /db/", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, http.StatusNotImplemented, utils.JSON{
			"result": "Not implemented!",
		})
	})
	mux.HandleFunc("GET /chaos/", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, http.StatusNotImplemented, utils.JSON{
			"result": "Not implemented!",
		})
	})
	mux.HandleFunc("GET /health/", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, http.StatusOK, utils.JSON{
			"result": "Ok",
		})
	})

	fmt.Println("Running on port", Env.Addr)
	log.Fatal(http.ListenAndServe(Env.Addr, mux))
}
