package database

import (
	"errors"
	"github.com/egasa21/hello-pet-api/config"
	"github.com/egasa21/hello-pet-api/infra/logger"
	"github.com/golang-migrate/migrate/v4"
	"os"
	"time"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

var (
	attempts = _defaultAttempts
	err      error
	m        *migrate.Migrate
)

func Migrate() {
	databaseURL, ok := os.LookupEnv(config.GetDNSConfig())
	if !ok || len(databaseURL) == 0 {
		logger.Fatal("Migrate: environment variable not declared: PG_URL")
	}

	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseURL)
		if err == nil {
			break
		}

		logger.Log("MigrateL postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		logger.Fatal("Migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if errors.Is(err, migrate.ErrNoChange) {
		logger.Log("Migrate: no change")
		return
	}

	logger.Log("Migrate: up success")
}
