package main

import (
	"C"
	"context"
	"github.com/joho/godotenv"
	"log"
	database "matcha/database/postgres"
	"matcha/pkg/routes"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Create base context
	ctx, cancel := context.WithCancel(context.Background())

	// Create database
	println(1111)
	db := database.Postgres(ctx) // todo:: used database.New()
	//db := database.TestDb(ctx) // for debug
	// Register endpoints
	println(1111)
	routes.Init(db)
	println(1111)

	// Sever
	err := http.ListenAndServe(os.Getenv("AUTH_PORT"), nil)
	if err != nil {
		cancel()
		log.Fatalln(err)
	}
}
