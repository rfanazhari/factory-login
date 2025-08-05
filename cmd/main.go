package main

import (
	di "github.com/rfanazhari/factory-login/internal/interfaces/http"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	// Load configuration from environment variables
	secretCaptcha := os.Getenv("GOOGLE_RECAPTCHA_SECRET")
	if secretCaptcha == "" {
		log.Fatal("GOOGLE_RECAPTCHA_SECRET environment variable is required")
	}

	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		redisUrl = "localhost:6379" // default value
	}

	maxRateLimitStr := os.Getenv("MAX_RATE_LIMIT")
	maxRateLimit := 5 // default value
	if maxRateLimitStr != "" {
		if parsed, err := strconv.Atoi(maxRateLimitStr); err == nil {
			maxRateLimit = parsed
		}
	}

	maxRateLimitDurationStr := os.Getenv("MAX_RATE_LIMIT_DURATION_MINUTES")
	maxRateLimitDuration := 15 * time.Minute // default value
	if maxRateLimitDurationStr != "" {
		if parsed, err := strconv.Atoi(maxRateLimitDurationStr); err == nil {
			maxRateLimitDuration = time.Duration(parsed) * time.Minute
		}
	}

	// Add the missing skipCaptcha parameter
	skipCaptcha := false // or read from environment variable
	if os.Getenv("SKIP_CAPTCHA") == "true" {
		skipCaptcha = true
	}

	container := di.NewContainer(secretCaptcha, redisUrl, maxRateLimit, maxRateLimitDuration, skipCaptcha)

	http.HandleFunc("/login", container.LoginHandler.Login)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
