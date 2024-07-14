package database

import (
	"github.com/egasa21/hello-pet-api/infra/logger"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"time"
)

type DB struct {
	Database *gorm.DB
}

func DBConnection(dsn string) (*DB, error) {
	logMode := viper.GetBool("MASTER_DB_LOG_MODE")
	logLevel := gormLog.Silent
	if logMode {
		logLevel = gormLog.Info
	}

	pgCon := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(pgCon, &gorm.Config{
		Logger: gormLog.Default.LogMode(logLevel),
	})

	if err != nil {
		logger.Fatal("database refused %v", err)
	}

	psqlDB, _ := db.DB()
	psqlDB.SetConnMaxIdleTime(time.Minute * 1)
	psqlDB.SetConnMaxLifetime(time.Minute * 5)

	err = psqlDB.Ping()
	if err != nil {
		logger.Fatal("database refused %v", err)
	}
	logger.Log("Database connection established")

	return &DB{Database: db}, nil
}
