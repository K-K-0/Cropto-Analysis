package hub

import "github.com/gorilla/websocket"

type Client struct {
	conn *websocket.Conn
	send chan []byte
	hub  *Hub
}

func NewClient(conn *websocket.Conn, hub *Hub) *Client {
	return &Client{
		conn: conn,
		send: make(chan []byte),
		hub:  hub,
	}
}

func (c *Client) WritePump() {
	defer c.conn.Close()
	for msg := range c.send {
		_ = c.conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.hub.Unregister(c)
		c.conn.Close()
	}()
	for {
		if _, _, err := c.conn.NextReader(); err != nil {
			break
		}
	}
}
