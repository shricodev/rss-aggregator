package initializers

import (
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DATABASE_CREDS")

	// Connect to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("There was an error connecting to the database")
	}

	sqlDB, _ := DB.DB()
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	log.Info().Msg("Connected to the database")
}
