package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal JSON response: %+v\n", payload)
		w.WriteHeader(500)
		return
	}

	//response header to the http request
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

// format the message into consistent message every single time
func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Println("responding with 5xx error", msg)
	}

	type ErrorResp struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, statusCode, ErrorResp{
		Error: msg,
	})
}
