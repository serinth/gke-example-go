package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()
	router.GET("/health", health)
	router.GET("/info", info)
	router.Run(":80")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func info(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "v1", "name": "Foo"})
}
