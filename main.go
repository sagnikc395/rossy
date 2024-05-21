package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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

	//setup headers
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://", "http://"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	// //connecting the handler readiness function
	// // to check is server is alive and running
	// v1Router.HandleFunc("/healthz", handlerReadiness)
	//on handle on get request
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

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
