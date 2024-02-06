package initializers

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func CheckEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("There was an error starting the server")
	}

	_, ok := os.LookupEnv("SERVER_PORT")

	if !ok {
		os.Setenv("SERVER_PORT", "8080")
	}

	_, ok = os.LookupEnv("DATABASE_CREDS")

	if !ok {
		log.Fatal().Msg("DATABASE_CREDS environment variable is not set")
	}
}
