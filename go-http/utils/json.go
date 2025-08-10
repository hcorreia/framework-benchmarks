package utils

import (
	"encoding/json"
	"fmt"
	"io"
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

func HttpGetJSON(url string, v any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err

	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}

	return err
}
