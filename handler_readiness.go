package main

import "net/http"

// function to define a http handler
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	//marshal with an empty json response
	respondWithJSON(w, 200, struct{}{})
}
