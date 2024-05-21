package main

import (
	"fmt"
	"log"
	"os"

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

	fmt.Printf("Port: %v\n", portString)
}
