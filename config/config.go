package config

import (
  "fmt"
  "os"

  "github.com/joho/godotenv"
  "github.com/spf13/cast"
)

type Config struct {
  HTTPPort string

  PostgresHost     string
  PostgresPort     int
  PostgresUser     string
  PostgresPassword string
  PostgresDatabase string

  DefaultOffset string
  DefaultLimit  string

}

func Load() Config {
  if err := godotenv.Load(); err != nil {
    fmt.Println("No .env file found")
  }

  config := Config{}

  config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7777"))

  config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "library_Song"))
  config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
  config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
  config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "20005"))
  config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "music_library"))

  config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
  config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))
  return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
  val, exists := os.LookupEnv(key)

  if exists {
    return val
  }

  return defaultValue
}
