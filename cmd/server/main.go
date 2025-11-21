package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThePratikSah/flag-zero/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := LoadConfig()

	router := Router()
	srv := HTTPServer(cfg, router)
	RunServer(srv)
}

// LoadConfig loads configuration from environment variables or .env file,
// validates it, and sets Gin mode if running in production.
// Panics or exits the process if config is invalid.
//
// Returns a pointer to a fully initialized config.Config struct for use
// throughout the application.
// Example:
//
//	cfg := LoadConfig()
//
// Returns:
//   - pointer to config.Config struct (with validated settings)
//
// Note: This helper abstracts config loading and validation in one place.
// If you need to customize config behavior, do so here.
func LoadConfig() *config.Config {
	cfg := config.LoadEnvConfig()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Failed to validate config: %v", err)
	}

	if cfg.App.Env == config.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	return cfg
}

// Router sets up the Gin HTTP router with all application routes and middleware.
// It defines endpoint handlers such as /health and /ready for basic health and readiness checks.
// Extend this function to add your application's API routes.
// Returns a fully configured *gin.Engine instance.
// Example:
//
//	router := Router()
func Router() *gin.Engine {
	router := gin.New()

	// middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// health and ready endpoints
	router.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	router.GET("/ready", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ready": true}) })

	// extend this function to add your application's API routes
	return router
}

// HTTPServer sets up and returns a new *http.Server instance configured with the
// provided settings from the config.Config struct and the given HTTP handler.
// It constructs the server's address from the configured host and port,
// and sets read/write timeouts as specified.
// This function is typically used to create the application's main HTTP server.
// Example:
//
//	srv := HTTPServer(cfg, router)
//	runServer(srv)
//
// Params:
//   - cfg:    pointer to loaded config.Config (with Server fields set)
//   - handler: main HTTP handler (e.g., a Gin engine)
//
// Returns:
//   - pointer to http.Server ready to be started by runServer()
//
// Note: This helper abstracts server construction details in one place.
// If you need to customize server behavior, do so here.
func HTTPServer(cfg *config.Config, handler http.Handler) *http.Server {
	addr := net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)
	return &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}
}

// RunServer starts the HTTP server and waits for it to shut down.
// It sets up signal handling for graceful shutdown and listens for incoming requests.
// Panics or exits the process if the server fails to start or shutdown.
//
// Example:
//
//	RunServer(srv)
//
// Params:
//   - srv: pointer to http.Server to be started
//
// Returns:
//   - none
//
// Note: This helper abstracts server startup and shutdown in one place.
// If you need to customize server behavior, do so here.
func RunServer(srv *http.Server) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
