package migrations

import (
	"github.com/egasa21/hello-pet-api/infra/database"
	"log"
)

func Migrate(db *database.DB) {
	var migrationModels []interface{}

	log.Println("Starting database migrations...")
	err := db.Database.AutoMigrate(migrationModels...)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migrations completed")
}
