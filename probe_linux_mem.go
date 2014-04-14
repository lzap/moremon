package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
	"math"
)

type MemLinuxResource struct {
}

func (r MemLinuxResource) Init() *InitPlotMesage {
	meminfo, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		panic("Cannot read memory info")
	}
	max := (math.Max(float64(meminfo["MemTotal"]), float64(meminfo["SwapTotal"])) * 1.20) / 10000
	return &InitPlotMesage{Id: "my", Type: "init", Label: []string{"used", "swap", "cache"}, Min: 0.0, Max: max}
}

func (r MemLinuxResource) Probe() (*UpdateMessage, error) {
	meminfo, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		panic("Cannot read memory info")
	}
	used := float64(meminfo["MemTotal"]-meminfo["MemFree"]-meminfo["Buffers"]-meminfo["Cached"]) / 10000
	swap := float64(meminfo["SwapTotal"]-meminfo["SwapFree"]) / 10000
	cached := float64(meminfo["Cached"]) / 10000
	return &UpdateMessage{Id: "my", V: []float64{used, swap, cached}}, nil
}
