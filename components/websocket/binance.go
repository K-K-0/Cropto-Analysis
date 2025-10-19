package websocket

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const baseUrl = "wss://stream.binance.com:9443/ws"

func StartStream(c context.Context, h *hub.Hub, symbol string) {
	stream := fmt.Sprintf("%s@trade", symbol)
	url := fmt.Sprintf("%s%s", baseUrl, stream)

	go func() {
		for {
			select {
			case <-c.Done():
				return
			default:
			}

			log.Printf("Connecting to Binance stream: %s", url)
			ws, _, err := websocket.DefaultDialer.Dial(url, nil)
			if err != nil {
				log.Println("Binance dial error:", err)
				time.Sleep(5 * time.Second)
				continue
			}
			for {
				_, msg, err := ws.ReadMessage()
				if err != nil {
					log.Println("Binance read error:", err)
					ws.Close()
					break
				}
				h.Broadcast(msg)
			}
			time.Sleep(5 * time.Second)
		}
	}()
}
