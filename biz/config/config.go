package config

type Config struct {
	Database *DatabaseConfig
	Jwt *JwtConfig
}

func New() *Config {
	return &Config{
		Database: NewDatabase(),
		Jwt: NewJwtConfig(),
	}
}