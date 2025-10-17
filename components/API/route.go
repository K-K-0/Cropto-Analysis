package api

import (
	"crypto-analysis/components/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/api/latest/:symbol", func(c *gin.Context) {
		symbol := c.Param("symbol")

		rows, err := database.DB.Query(`SELECT price, quantity, timestamp FROM 
					trades WHERE symbol=$1 ORDER BY id DESC LIMIT 50`, symbol)

		if err != nil {
			log.Println("Query error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		defer rows.Close()

		var data []map[string]interface{}

		for rows.Next() {
			var price, qty string
			var ts string
			rows.Scan(&price, &qty, &ts)
			data = append(data, gin.H{"Price": price, "Quantity": qty, "Timestamp": ts})
		}
		c.JSON(http.StatusOK, gin.H{"symbol": symbol, "trades": data})
	})
}
