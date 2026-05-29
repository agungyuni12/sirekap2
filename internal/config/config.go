package config

import (
	"log"
	"os"
	"strings"
)

// Config holds the application configuration
type Config struct {
	DB      DBConfig
	Server  ServerConfig
	Storage StorageConfig
}

// DBConfig holds database configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Addr          string
	SessionKey    string
	SessionSecure bool
}

// StorageConfig holds object storage configuration.
type StorageConfig struct {
	EndpointURL     string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	Region          string
	UsePathStyle    bool
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvFirst(keys ...string) string {
	for _, key := range keys {
		if value, exists := os.LookupEnv(key); exists && strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func getEnvBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "true", "yes", "y", "on":
		return true
	case "0", "false", "no", "n", "off":
		return false
	default:
		return defaultValue
	}
}

// Load returns the application configuration
func Load() *Config {
	serverAddr := strings.TrimSpace(getEnvFirst("SERVER_ADDR"))
	if serverAddr == "" {
		port := strings.TrimSpace(getEnvFirst("PORT"))
		if port != "" {
			serverAddr = "0.0.0.0:" + port
		} else {
			serverAddr = "0.0.0.0:8051"
		}
	}

	dbPassword := strings.TrimSpace(getEnvFirst("DB_PASSWORD", "DB_PASS"))
	if dbPassword == "" {
		log.Fatal("FATAL: DB_PASSWORD atau DB_PASS harus diset via environment variable")
	}

	sessionKey := strings.TrimSpace(getEnvOrDefault("SESSION_KEY", ""))
	if len(sessionKey) < 32 {
		log.Fatal("FATAL: SESSION_KEY harus diset via environment variable (minimal 32 karakter)")
	}

	return &Config{
		DB: DBConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvOrDefault("DB_PORT", "3306"),
			User:     getEnvOrDefault("DB_USER", "root"),
			Password: dbPassword,
			DBName:   getEnvOrDefault("DB_NAME", "seleksimitra"),
		},
		Server: ServerConfig{
			Addr:         serverAddr,
			SessionKey:   sessionKey,
			SessionSecure: getEnvBool("SESSION_SECURE", true),
		},
		Storage: StorageConfig{
			EndpointURL:     strings.TrimSpace(getEnvFirst("AWS_ENDPOINT_URL")),
			AccessKeyID:     strings.TrimSpace(getEnvFirst("AWS_ACCESS_KEY_ID")),
			SecretAccessKey: strings.TrimSpace(getEnvFirst("AWS_SECRET_ACCESS_KEY")),
			Bucket:          strings.TrimSpace(getEnvFirst("AWS_BUCKET")),
			Region:          firstNonEmpty(strings.TrimSpace(getEnvFirst("AWS_DEFAULT_REGION")), "us-east-1"),
			UsePathStyle:    getEnvBool("AWS_USE_PATH_STYLE_ENDPOINT", true),
		},
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}
