package websocket

import (
	"crypto-analysis/components/database"
	"crypto-analysis/components/model"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type BinanceSub struct {
	Type   string   `json:"type"`
	Params []string `json:"params"`
	ID     int      `json:"id"`
}

func StartBinanceStream(symbols []string) {
	url := "wss://stream.binance.com:9443/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error while establishing connection with binance websocket: ", err)
	}

	defer conn.Close()
	log.Println("connected binance websocket connection")

	for _, s := range symbols {
		sub := map[string]interface{}{
			"method": "SUBSCRIBE",
			"params": []string{s + "@trade"},
			"id":     time.Now().Unix(),
		}
		if err := conn.WriteJSON(sub); err != nil {
			log.Println("subscribe error:", err)
		}
	}

	go func() {
		for {
			err := conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Println("ping error", err)
				return
			}
			time.Sleep(30 * time.Second)
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("error while reading:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		var trade struct {
			EventType string `json:"e"`
			Symbol    string `json:"s"`
			Price     string `json:"p"`
			Quantity  string `json:"q"`
		}

		if err := json.Unmarshal(msg, &trade); err != nil {
			continue
		}

		if trade.EventType != "trade" {
			continue
		}

		data := model.Trade{
			Symbol:    trade.Symbol,
			Price:     trade.Price,
			Quantity:  trade.Quantity,
			Timestamp: time.Now(),
		}

		_, err = database.DB.Exec(`INSERT INTO trades (symbol, price, quantity, timestamp) VALUES ($1,$2,$3,$4)`,
			data.Symbol, data.Price, data.Quantity, data.Timestamp)
		if err != nil {
			log.Println("Database insert Error: ", err)
		}
		log.Printf("[%s] %s @ %s", trade.Symbol, trade.Quantity, trade.Price)

	}
}
