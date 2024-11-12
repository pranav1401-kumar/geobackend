package config

import (
    "os"
)

type Config struct {
    JWTSecret string
    DBHost    string
    DBUser    string
    DBName    string
    DBPass    string
    DBPort    string
}

func LoadConfig() Config {
    return Config{
        JWTSecret: os.Getenv("JWT_SECRET"),
        DBHost:    os.Getenv("DB_HOST"),
        DBUser:    os.Getenv("DB_USER"),
        DBName:    os.Getenv("DB_NAME"),
        DBPass:    os.Getenv("DB_PASS"),
        DBPort:    os.Getenv("DB_PORT"),
    }
}
