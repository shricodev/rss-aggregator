package config

import {
  "golang-prisma/prisma/db"
  "github.com/rs/zerolog/log"
}

funct connectDB(*db.PrismaClient, error) {
  client := db.NewClient()

  if err := client.Prisma.Connect(); err != nil {
    return nil, err
  }

  log.Info().Msg("Connected to the database")
  return client, nil
}
