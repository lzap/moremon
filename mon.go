package main

import (
	"time"
)

type Resource interface {
	Init() *InitPlotMesage
	Probe() (*UpdateMessage, error)
}

type Monitor struct {
	updateTicker *time.Ticker
	history int
	resources map[string]Resource
}

var monitor = NewMonitor()

func NewMonitor() *Monitor {
	res := map[string]Resource {
		"ld": LoadLinuxResource{},
	}
	return &Monitor{ resources: res }
}

func (m *Monitor) setUpdateInterval(interval int) {
	if m.updateTicker != nil {
		m.updateTicker.Stop()
	}
	m.updateTicker = time.NewTicker(time.Duration(interval) * time.Second)
}

func (m *Monitor) setHistorySize(size int) {
	m.history = size
}

func (m *Monitor) readData() {
	for {
		select {
		case <-m.updateTicker.C:
			for name, resource := range m.resources {
				msg, err := resource.Probe()
				if err != nil {
					eLogger.Println("Unable to probe", name, "-", err)
					continue
				}
				h.broadcast <- msg.asJSON()
			}
		}
	}
}
