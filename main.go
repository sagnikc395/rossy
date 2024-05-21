package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	//load PORT from .env files s
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		//exit the program and return the port not bound in env
		log.Fatal("PORT not found in the environment")
	}

	//fmt.Printf("Port: %v\n", portString)

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v\n", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
