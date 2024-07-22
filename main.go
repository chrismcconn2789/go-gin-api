package main

import (
	"time"

	"github.com/chrismconn2789/go-gin-api/db"
	"github.com/chrismconn2789/go-gin-api/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Name   string `json:"name"`
	Amount uint   `json:"amount"`
}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true
    config.MaxAge = 12 * time.Hour
	router.Use(cors.Default())

	db.ConnectDatabase()

	router.POST("/transactions", handlers.CreateTransaction)
	router.GET("/transactions", handlers.FindTransactions)
	router.GET("/transactions/:id", handlers.FindTransactionById)
	router.DELETE("/transactions/:id", handlers.DeleteTransaction)

	router.Run()
}
