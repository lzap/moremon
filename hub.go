package main

import (
	"time"
)

const (
	readInterval = time.Second
)

type Hub struct {
	connections map[*Connection]bool
	broadcast   chan []byte
	register    chan *Connection
	unregister  chan *Connection
}

var h = Hub{
	broadcast:   make(chan []byte),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
	connections: make(map[*Connection]bool),
}

func (h *Hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
			dLogger.Println("Registered connection from %s", c.fromAddr)
			monitor.SendInitMessages(c)
		case c := <-h.unregister:
			delete(h.connections, c)
			close(c.send)
			dLogger.Println("Unregistered connection from %s", c.fromAddr)
		case m := <-h.broadcast:
			for c := range h.connections {
				dLogger.Println("Sending", string(m), "to", c.fromAddr)
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(h.connections, c)
				}
			}
		}
	}
}
