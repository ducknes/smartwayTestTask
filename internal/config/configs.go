package config

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
		Host:        DBHOST,
		Port:        "5432",
		Username:    DBUSERNAME,
		Password:    DBPASSWORD,
		Database:    DBNAME,
		MaxAttempts: 5,
	}
}
