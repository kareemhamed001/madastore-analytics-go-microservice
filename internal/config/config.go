package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	Environment string
	ServerPort  string
	DatabaseDSN string
	ApiKey      string
	GRPCPort    string

	RedisAddr     string
	RedisPassword string
	RedisDB       int

	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBConnMaxLifetime time.Duration

	HTTPReadTimeout       time.Duration
	HTTPWriteTimeout      time.Duration
	HTTPIdleTimeout       time.Duration
	HTTPReadHeaderTimeout time.Duration
	RequestTimeout        time.Duration

	ShutdownTimeout      time.Duration
	CacheRefreshInterval time.Duration
	RepoQueryTimeout     time.Duration
}

func Load() *Config {
	_ = godotenv.Load()
	return &Config{
		AppName:     getEnvMulti([]string{"APP_NAME", "AppName"}, "analytics"),
		Environment: getEnvMulti([]string{"ENVIRONMENT", "Environment"}, "development"),
		ServerPort:  getEnvMulti([]string{"SERVER_PORT", "ServerPort"}, "8080"),
		DatabaseDSN: getEnvMulti([]string{"DATABASE_DSN", "DatabaseDSN"}, ""),
		ApiKey:      getEnv("API_KEY", ""),
		GRPCPort:    getEnv("GRPC_PORT", "9090"),

		RedisAddr:     getEnv("REDIS_ADDR", "127.0.0.1:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),

		DBMaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 25),
		DBMaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 25),
		DBConnMaxLifetime: getEnvDuration("DB_CONN_MAX_LIFETIME", 30*time.Minute),

		HTTPReadTimeout:       getEnvDuration("HTTP_READ_TIMEOUT", 30*time.Second),
		HTTPWriteTimeout:      getEnvDuration("HTTP_WRITE_TIMEOUT", 30*time.Second),
		HTTPIdleTimeout:       getEnvDuration("HTTP_IDLE_TIMEOUT", 60*time.Second),
		HTTPReadHeaderTimeout: getEnvDuration("HTTP_READ_HEADER_TIMEOUT", 5*time.Second),
		RequestTimeout:        getEnvDuration("REQUEST_TIMEOUT", 20*time.Second),

		ShutdownTimeout:      getEnvDuration("SHUTDOWN_TIMEOUT", 10*time.Second),
		CacheRefreshInterval: getEnvDuration("CACHE_REFRESH_INTERVAL", 5*time.Minute),
		RepoQueryTimeout:     getEnvDuration("REPO_QUERY_TIMEOUT", 3*time.Second),
	}
}

func getEnv(key, falllback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return falllback
}

func getEnvMulti(keys []string, fallback string) string {
	for _, key := range keys {
		if value, exists := os.LookupEnv(key); exists {
			return value
		}
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	parsed, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return parsed
}
