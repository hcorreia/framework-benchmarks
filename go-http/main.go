package main

import (
	"fmt"
	"framework-benchmarks/go-http/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting hostname: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, http.StatusOK, utils.JSON{
			"result":   "Ok",
			"hostname": hostname,
		})
	})
	mux.HandleFunc("GET /db/", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, http.StatusNotImplemented, utils.JSON{
			"result":   "Not implemented!",
			"hostname": hostname,
		})
	})
	mux.HandleFunc("GET /chaos/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			SleepTime int `json:"sleep_time"`
		}{}

		err = utils.HttpGetJSON(Env.ChaosEndpoint, &data)
		if err != nil {
			utils.WriteJSON(w, http.StatusServiceUnavailable, utils.JSON{
				"result":   "Chaos service is unavailable.",
				"hostname": hostname,
			})
			log.Printf("Chaos service is unavailable: %v", err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, utils.JSON{
			"result":     "Ok",
			"hostname":   hostname,
			"sleep_time": data.SleepTime,
		})
	})
	mux.HandleFunc("GET /health/", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, http.StatusOK, utils.JSON{
			"result":   "Ok",
			"hostname": hostname,
		})
	})

	fmt.Println("Running on port", Env.Addr)
	log.Fatal(http.ListenAndServe(Env.Addr, mux))
}
