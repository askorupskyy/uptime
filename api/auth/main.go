package main

import (
	"context"
	gin "github.com/gin-gonic/gin"
	"github.com/rcbxd/uptime/api/auth/handlers"
	"github.com/rcbxd/uptime/api/auth/middleware"
	"github.com/rcbxd/uptime/api/auth/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.JSONLogMiddleware(utils.NewLogger()))
	r.POST("/login", handlers.LoginHandler)

	// health check endpoint
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping-pong",
		})
	})

	// initialize the server and use our gin router as a handler
	srv := &http.Server{
		// TODO: make this configurable through docker compose and .env
		Addr:    ":9000",
		Handler: r,
	}

  // non-blocking goroutine to start the server
	go func() {
		srv.ListenAndServe()
	}()

	// create signals for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// wait for signal
	<-sigs

	// wait for all background tasks to terminate and then terminate the server
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(tc)
}
