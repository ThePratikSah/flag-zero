package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	cfg  *Config
	once sync.Once
)

// LoadEnvConfig loads the config from the environment variables.
func LoadEnvConfig() *Config {
	// once.Do(func() { ... }) is a Go concurrency pattern that guarantees a piece of code runs only once,
	// no matter how many times it is called, even across multiple goroutines.
	once.Do(func() {
		// Load the environment variables from the .env file.
		err := godotenv.Load()
		if err != nil {
			log.Printf("[WARN] .env not loaded: %v", err)
		}

		// Initialize the config struct.
		cfg = &Config{
			App: AppConfig{
				Env:      Environment(getEnv("APP_ENV", "dev")),
				LogLevel: getEnv("APP_LOG_LEVEL", "info"),
			},
			Server: ServerConfig{
				Host:         getEnv("SERVER_HOST", "0.0.0.0"),
				Port:         getEnv("SERVER_PORT", "8080"),
				ReadTimeout:  getInt("SERVER_READ_TIMEOUT", 5),
				WriteTimeout: getInt("SERVER_WRITE_TIMEOUT", 10),
			},
			Database: DatabaseConfig{
				MySQLDSN: getEnv("MYSQL_DSN", ""),
				Host:     getEnv("DB_HOST", "localhost"),
				Port:     getEnv("DB_PORT", "3306"),
				User:     getEnv("DB_USER", "root"),
				Password: getEnv("DB_PASSWORD", "password"),
				Name:     getEnv("DB_NAME", "flag_zero"),
			},
			Redis: RedisConfig{
				Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
				Password: getEnv("REDIS_PASSWORD", ""),
				DB:       getInt("REDIS_DB", 0),
			},
		}
	})

	return cfg
}

// Check panics if the error is not nil.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// getEnv returns the value of the environment variable or fallback.
func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

// getInt converts env variable to int with fallback.
func getInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("[WARN] invalid int for %s: %v, using fallback", key, err)
		return fallback
	}
	return i
}

func (c *Config) Validate() error {
	if c.App.Env == "prod" {
		c.App.LogLevel = "error" // set log level to error in production
		if c.Database.MySQLDSN == "" {
			panic("MYSQL_DSN is required in production") // panic if MYSQL_DSN is not set
		}
		if c.Redis.Addr == "" {
			panic("REDIS_ADDR is required in production") // panic if REDIS_ADDR is not set
		}
		if c.Redis.Password == "" {
			panic("REDIS_PASSWORD is required in production") // panic if REDIS_PASSWORD is not set
		}
		if c.Redis.DB == 0 {
			panic("REDIS_DB is required in production") // panic if REDIS_DB is not set
		}
	}
	return nil
}
