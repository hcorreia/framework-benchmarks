package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JSON = map[string]any

func WriteJSON(w http.ResponseWriter, statusCode int, v JSON) error {
	result, err := json.Marshal(v)

	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error")

		return err
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)

	return nil
}
