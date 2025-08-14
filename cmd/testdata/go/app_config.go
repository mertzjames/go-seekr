// app_config.go
// Go application configuration with embedded secrets

package main

import (
	"os"
)

type Config struct {
	// GitHub API token for repository access
	GitHubToken string

	// Docker registry credentials
	DockerUser     string
	DockerPassword string

	// JWT secret key
	JWTSecret string
}

func LoadConfig() *Config {
	return &Config{
		// WARNING: Never hardcode tokens in production!
		GitHubToken:    "ghp_1234567890abcdefghijklmnopqrstuvwxyz123",
		DockerUser:     "mycompany-docker",
		DockerPassword: "docker-secret-2024!",
		JWTSecret:      "super-secret-jwt-key-change-in-prod",
	}
}

// Example of environment variable that might contain secrets
var (
	GOOGLE_CLIENT_SECRET = "GOCSPX-AbCdEfGhIjKlMnOpQrStUvWxYz123456"
	FIREBASE_API_KEY     = "AIzaSyBcdefghijklmnopqrstuvwxyz1234567890"
	SLACK_BOT_TOKEN      = "xoxb-1234567890123-4567890123456-AbCdEfGhIjKlMnOpQrStUvWx"
)

func main() {
	// Reading sensitive env vars (bad practice if hardcoded)
	dbPassword := os.Getenv("DATABASE_PASSWORD") // "prod_db_pass_2024!"
	apiKey := os.Getenv("STRIPE_API_KEY")        // "sk_live_51ExampleKey123456789"
}
