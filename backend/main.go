package main

import (
	"log"
	"os"

	"vehicle-telemetry-system/backend/database"
	"vehicle-telemetry-system/backend/handlers"
	"vehicle-telemetry-system/backend/routes"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "telemetry.db"
	}

	db, err := database.Open(dbPath)
	if err != nil {
		log.Fatalf("database open failed: %v", err)
	}
	defer db.Close()

	telemetryHandler := handlers.NewTelemetryHandler(db)
	router := routes.SetupRouter(telemetryHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Telemetry API listening on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
