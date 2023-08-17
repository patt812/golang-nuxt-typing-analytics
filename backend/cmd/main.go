package main

import (
	"log"

	"github.com/patt812/golang-nuxt-typing-analytics/api"
	"github.com/patt812/golang-nuxt-typing-analytics/database"
)

func main() {
	db, err := database.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return
	}

	r := api.Router(db)
	r.Run(":8080")
}
