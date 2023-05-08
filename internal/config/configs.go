package config

import (
	"os"
)

type StorageConfig struct {
	Host        string
	Port        string
	Username    string
	Password    string
	Database    string
	MaxAttempts int
}

func NewStorageConfig() StorageConfig {
	return StorageConfig{
		Host:        os.Getenv("DBHOST"),
		Port:        "5432",
		Username:    os.Getenv("DBUSERNAME"),
		Password:    os.Getenv("DBPASSWORD"),
		Database:    os.Getenv("DBNAME"),
		MaxAttempts: 5,
	}
}
