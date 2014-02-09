package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	maximumAllowedTime = 24 * time.Hour
	maxMessageSize = 512
)

type Connection struct {
	fromAddr string
	ws *websocket.Conn
	send chan []byte
}

func (c *Connection) readPump() {
	defer func() {
		h.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(maximumAllowedTime))
	for {
		_, _, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		// we do not expect any input
	}
}

func (c *Connection) writePump() {
	defer func() {
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.ws.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	// we do not check origin - it's a feature :-)
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	c := &Connection{send: make(chan []byte, 256), ws: ws, fromAddr: r.RemoteAddr}
	h.register <- c
	go c.writePump()
	c.readPump()
}
