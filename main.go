package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"

	"github.com/shricodev/rss-aggregator/handlers"
	"github.com/shricodev/rss-aggregator/initilizers"
)

func init() {
	initilizers.CheckEnvVariables()
	initilizers.ConnectToDB()
}

func main() {
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

	v1Router.Get("/healthz", handlers.HandlerReadiness)
	v1Router.Get("/err", handlers.HandlerError)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + serverPort,
	}

	fmt.Printf("The server is listening on the port: %s", serverPort)
	serverError := server.ListenAndServe()

	// This is only executed if the server ever throws an error.
	if serverError != nil {
		log.Fatal().Err(serverError).Msg("There was an error starting the server")
	}
}
