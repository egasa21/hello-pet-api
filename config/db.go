package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Driver      string
	DBName      string
	Username    string
	Password    string
	Host        string
	Port        string
	LogMode     string
	ClusterName string // Added for CockroachDB cluster name
}

func GetDNSConfig() string {
	masterName := viper.GetString("MASTER_DB_NAME")
	masterUser := viper.GetString("MASTER_DB_USER")
	masterPassword := viper.GetString("MASTER_DB_PASSWORD")
	masterHost := viper.GetString("MASTER_DB_HOST")
	masterPort := viper.GetString("MASTER_DB_PORT")
	masterSslMode := viper.GetString("MASTER_SSL_MODE")

	// CockroachDB connection string format
	masterDSN := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		masterUser, masterPassword, masterHost, masterPort, masterName, masterSslMode,
	)
	return masterDSN
}
