package main

import (
	"crypto-analysis/components/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	// go websocket.StartBinanceStream([]string{"btcusdt", "ethusdt", "solusdt"})

	r := gin.Default()
	// api.SetupRoutes(r)
	r.Run(":8080")
}
