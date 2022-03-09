package config

import (
	"fmt"
)

// DBConfig represents db configuration
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	DBName   string `yaml:"name"`
	Password string `yaml:"password"`
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable&statement_cache_mode=describe",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
