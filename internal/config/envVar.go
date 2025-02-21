package config

import (
	// "github.com/joho/godotenv"
	// "log"
	"os"
)

type Config struct {
	Port               string
	MongoURI           string
	MongoDatabase           string
	AccessTokenSecret  string
	RefreshTokenSecret string
	TMOBAPIKey         string
}

func LoadConfig() *Config {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("err loading: %v", err)
	// }
	return &Config{
		Port:               getEnv("PORT", "3000"),
		MongoURI:           getEnv("MONGO_URI", ""),
		MongoDatabase:      getEnv("MONGO_DATABASE", ""),
		AccessTokenSecret:  getEnv("ACCESS_TOKEN_SECRET", ""),
		RefreshTokenSecret: getEnv("REFRESH_TOKEN_SECRET", ""),
		TMOBAPIKey:         getEnv("TMOB_API_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
