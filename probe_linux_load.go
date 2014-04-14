package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
)

type LoadLinuxResource struct {
}

func numCores() int {
	cpuinfo, err := linuxproc.ReadCPUInfo("/proc/cpuinfo")
	if err != nil {
		return 16
	}
	return cpuinfo.NumCPU() * 2
}

func (r LoadLinuxResource) Init() *InitPlotMesage {
	return &InitPlotMesage{Id: "ld", Type: "init", Label: []string{"load"}, Min: 0.0, Max: float64(numCores())}
}

func (r LoadLinuxResource) Probe() (*UpdateMessage, error) {
	loadavg, err := linuxproc.ReadLoadAvg("/proc/loadavg")
	if err != nil {
		return nil, err
	}
	return &UpdateMessage{Id: "ld", V: []float64{loadavg.Last1Min}}, nil
}
