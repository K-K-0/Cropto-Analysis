package main

import (
	api "crypto-analysis/components/API"
	"crypto-analysis/components/database"
	"crypto-analysis/components/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	go websocket.StartBinanceStream([]string{"btcusdt", "ethusdt", "solusdt"})

	r := gin.Default()
	api.SetupRoutes(r)
	r.Run(":8080")
}
