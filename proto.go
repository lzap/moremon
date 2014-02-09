package main

import (
	"encoding/json"
)

type InitPlotMesage struct {
	Id string
	Label string
	Min float64
	Max float64
}

type UpdateMessage struct {
	Id string
	V float64
}

func (m *UpdateMessage) asJSON() []byte {
	result, err := json.Marshal(m)
	if err != nil {
		eLogger.Panic("Unable to marshall JSON for", m)
	}
	dLogger.Println("Marshalled", m, "as", string(result))
	return result
}
