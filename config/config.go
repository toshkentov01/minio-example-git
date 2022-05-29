package config

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	BucketName string
	Location   string
	EndPoint   string
	AccessID   string
	SecretKey  string
}

func load() *Config {
	return &Config{
		BucketName: cast.ToString(getOrReturnDefault("BUCKET_NAME", "")),
		Location:   cast.ToString(getOrReturnDefault("LOCATION", "")),
		EndPoint:   cast.ToString(getOrReturnDefault("END_POINT", "")),
		AccessID:   cast.ToString(getOrReturnDefault("ACCESS_ID", "")),
		SecretKey:  cast.ToString(getOrReturnDefault("SECRET_KEY", "")),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

var (
	instance *Config
	once     sync.Once
)

//Get ...
func Get() *Config {
	once.Do(func() {
		instance = load()
	})

	return instance
}
