package main

import (
	"github.com/shricodev/rss-aggregator/initializers"
	"github.com/shricodev/rss-aggregator/models"
)

func init() {
	initializers.CheckEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
