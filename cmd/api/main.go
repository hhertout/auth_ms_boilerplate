package main

import (
	"auth_ms/migrations"
	"auth_ms/src/router"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if os.Getenv("DOCKER") != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	m := migrations.NewMigration("/")
	if err := m.MigrateAll(); err != nil {
		log.Println(err)
		return
	}

	r := router.Serve()
	log.Printf("📡 Server start on port %s \n", os.Getenv("PORT"))
	if err := r.Run(); err != nil {
		fmt.Println("Error on running server")
		fmt.Printf("Error: %s", err)
		return
	}
}
