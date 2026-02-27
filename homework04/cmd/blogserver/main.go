package main

import (
	"log"
	"os"

	"homework04/internal/blog"
)

func getenvWithDefault(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func main() {
	dsn := getenvWithDefault("BLOG_DSN", "blog.db")
	secret := getenvWithDefault("BLOG_JWT_SECRET", "homework04-dev-secret")
	addr := getenvWithDefault("BLOG_ADDR", ":8080")

	logger := log.New(os.Stdout, "[blogserver] ", log.LstdFlags)
	app, err := blog.NewSQLiteApp(blog.Config{
		DSN:       dsn,
		JWTSecret: secret,
		Logger:    logger,
	})
	if err != nil {
		logger.Fatalf("init app failed: %v", err)
	}

	logger.Printf("server listening on %s (dsn=%s)", addr, dsn)
	if err := app.Router.Run(addr); err != nil {
		logger.Fatalf("server stopped: %v", err)
	}
}
