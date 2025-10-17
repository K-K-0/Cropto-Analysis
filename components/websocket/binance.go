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
	log.Println("connected binance websocket connection")

	sub := BinanceSub{
		Type:   "SUBSCRIBE",
		Params: []string{},
		ID:     1,
	}

	for _, s := range symbols {
		sub.Params = append(sub.Params, s+"@trade")
	}

	subMsg, _ := json.Marshal(sub)
	if err := conn.WriteMessage(websocket.TextMessage, subMsg); err != nil {
		log.Println("error while subscribe: ", err)
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("error while reading:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		var data map[string]interface{}

		if err := json.Unmarshal(msg, &data); err != nil {
			continue
		}

		if data["e"] != "trade" {
			continue
		}

		trade := model.Trade{
			Symbol:    data["symbol"].(string),
			Price:     data["price"].(string),
			Quantity:  data["quantity"].(string),
			Timestamp: time.Now(),
		}

		_, err = database.DB.Exec(`INSERT INTO trades (symbol, price, quantity) VALUES ($1,$2,$3,$4)`,
			trade.Symbol, trade.Price, trade.Quantity, trade.Timestamp)
		if err != nil {
			log.Println("Database insert Error: ", err)
		}
		log.Printf("[%s] %s @ %s", trade.Symbol, trade.Quantity, trade.Price)

	}
}
