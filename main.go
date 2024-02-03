package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Try to read the .env file in the current directory.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, ok := os.LookupEnv("SERVER_PORT")

	if !ok {
		os.Setenv("SERVER_PORT", "8080")
	}

	serverPort := os.Getenv("SERVER_PORT")
	fmt.Println("Server port is: " + serverPort)

	router := chi.NewRouter()
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}

	corsHandler := cors.New(corsOptions).Handler
	router.Use(corsHandler)

	v1Router := chi.NewRouter()

	router.Mount("/v1", v1Router)

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + serverPort,
	}

	fmt.Printf("The server is listening on the port: %s", serverPort)
	serverError := server.ListenAndServe()

	// This is only executed if the server ever throws an error.
	if serverError != nil {
		log.Fatal(err)
	}
}
