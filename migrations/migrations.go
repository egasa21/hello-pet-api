package migrations

import (
	"github.com/egasa21/hello-pet-api/infra/database"
	"github.com/egasa21/hello-pet-api/models/customer_model"
	"github.com/egasa21/hello-pet-api/models/doctor_model"
	"github.com/egasa21/hello-pet-api/models/user_model"
	"log"
)

func Migrate(db *database.DB) {
	var migrationModels = []interface{}{&user_model.User{}, customer_model.Customer{}, doctor_model.Doctor{}}

	log.Println("Starting database migrations...")
	err := db.Database.AutoMigrate(migrationModels...)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migrations completed")
}
