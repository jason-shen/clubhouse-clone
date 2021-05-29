package config

import "github.com/jason-shen/clubhouse-clone-biz/utils"

type DatabaseConfig struct {
	Host     string
	Name     string
	Password string
	User     string
	Port     string
}

func NewDatabase() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     utils.GetIni("database", "HOST", "localhost"),
		Name:     utils.GetIni("database", "DATABASE_NAME", "clubhouseclone_db"),
		User:     utils.GetIni("database", "DATABASE_USER", "postgres"),
		Password: utils.GetIni("database", "DATABASE_PASSWORD", "postgres"),
		Port:     utils.GetIni("database", "DATABASE_PORT", "5432"),
	}
}
