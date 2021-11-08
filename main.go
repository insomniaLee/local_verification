package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":7777") // start service with port 7777 to celebrate EDG win S11
	if err != nil {
		log.Printf("service start failed ,err inso: %v", err)
	}
}
