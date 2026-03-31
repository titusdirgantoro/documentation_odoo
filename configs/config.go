package configs

import (
	"crypto/sha256"
	"encoding/hex"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Config holds all runtime configuration for the server.
type Config struct {
	Port           string
	GinMode        string
	ImpactPassword string
	OwnerPassword  string
	CookieSecret   string
	ImpactToken    string
	OwnerToken     string
}

// Load reads environment variables and returns a populated Config.
// Falls back to sensible defaults when variables are not set.
func Load() *Config {
	// Attempt to load .env file; it will fail silently if not found,
	// which is fine since we fallback to OS environment variables.
	_ = godotenv.Load()

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	impactPwd := os.Getenv("IMPACT_PASSWORD")
	if impactPwd == "" {
		impactPwd = "impact_secret"
	}

	ownerPwd := os.Getenv("OWNER_PASSWORD")
	if ownerPwd == "" {
		ownerPwd = "owner_secret"
	}

	cookieSecret := os.Getenv("COOKIE_SECRET")
	if cookieSecret == "" {
		cookieSecret = "default-insecure-cookie-key"
	}

	impactHash := sha256.Sum256([]byte(cookieSecret + "_impact_salt"))
	ownerHash := sha256.Sum256([]byte(cookieSecret + "_owner_salt"))

	return &Config{
		Port:           port,
		GinMode:        mode,
		ImpactPassword: impactPwd,
		OwnerPassword:  ownerPwd,
		CookieSecret:   cookieSecret,
		ImpactToken:    hex.EncodeToString(impactHash[:]),
		OwnerToken:     hex.EncodeToString(ownerHash[:]),
	}
}
