package config

import (
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    DBType    string
    DBHost    string
    DBPort    int
    DBUser    string
    DBPass    string
    DBName    string
    ServerPort string
    LogLevel  string
}

func LoadConfig() (*Config, error) {
    // Load .env file, ignore error if already loaded in env
    _ = godotenv.Load()

    port, err := strconv.Atoi(getEnv("DB_PORT", "3306"))
    if err != nil {
        return nil, err
    }

    cfg := &Config{
        DBType:     getEnv("DB_TYPE", "mysql"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     port,
        DBUser:     getEnv("DB_USER", ""),
        DBPass:     getEnv("DB_PASS", ""),
        DBName:     getEnv("DB_NAME", "abc_user_db"),
        ServerPort: getEnv("SERVER_PORT", "8080"),
        LogLevel:   getEnv("LOG_LEVEL", "info"),
    }

    return cfg, nil
}

func getEnv(key, fallback string) string {
    if val := os.Getenv(key); val != "" {
        return val
    }
    return fallback
}

