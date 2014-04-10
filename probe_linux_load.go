package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
)

const (
	probeId    = "ld"
	probeLabel = "load"
)

type LoadLinuxResource struct {
}

func (r LoadLinuxResource) Init() *InitPlotMesage {
	return &InitPlotMesage{Id: probeId, Type: "init", Label: probeLabel, Min: 0.0, Max: 3.0}
}

func (r LoadLinuxResource) Probe() (*UpdateMessage, error) {
	loadavg, err := linuxproc.ReadLoadAvg("/proc/loadavg")
	if err != nil {
		return nil, err
	}
	return &UpdateMessage{Id: probeId, V: loadavg.Last1Min}, nil
}
