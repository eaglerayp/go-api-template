package apiserver

import (
	"context"
	"log"
	"net/http"

	"github.com/eaglerayp/go-api-template/gat"
	ginprometheus "github.com/eaglerayp/go-gin-prometheus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitGinServer(ctx context.Context) *http.Server {
	// Create gin http server.
	gin.SetMode(viper.GetString("http.mode"))
	router := gin.New()
	router.Use(gin.Recovery())

	// Add gin prometheus metrics
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	// general service for debugging
	router.GET("/health", health)
	router.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gat.GetVersion())
	})
	router.GET("/config", appConfig)
	router.GET("/mongo", mongoInfo)
	router.GET("/ready", ready)

	// TODO: Add application routing

	port := viper.GetString("http.port")
	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  viper.GetDuration("http.read_timeout"),
		WriteTimeout: viper.GetDuration("http.write_timeout"),
	}

	go func() {
		log.Printf("Server is running and listening port: %s\n", port)
		// service connections
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return httpServer
}
func mongoInfo(c *gin.Context) {
	status := map[string]interface{}{}
	c.JSON(http.StatusOK, status)
}

func appConfig(c *gin.Context) {
	settings := viper.AllSettings()
	delete(settings, "database")
	c.JSON(http.StatusOK, settings)
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func ready(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
