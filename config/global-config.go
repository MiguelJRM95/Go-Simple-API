package config

import(
	"os"
)

type Config struct{
	DataBaseConfig DataBaseConfig
}

func New() *Config {
	return &Config{
		DataBaseConfig: DataBaseConfig{
			Host: getEnv("DB_HOST", ""),
			User: getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Port: getEnv("DB_PORT", ""),
			Name: getEnv("DB_NAME", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
		return value
    }

    return defaultVal
}