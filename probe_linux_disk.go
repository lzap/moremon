package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
	"regexp"
	"strings"
)

type DiskLinuxResource struct {
	MountPoints []string
}

func NewDiskLinuxResource() *DiskLinuxResource {
	validFS := regexp.MustCompile(`^(cifs|btrfs|ext\d?|hfs|jfs|minix|nfs\d?|ntfs|reiserfs|smbfs|vfat|xfs)$`)
	mounts, err := linuxproc.ReadMounts("/proc/mounts")
	if err != nil {
		panic("Cannot read mounts info")
	}
	mountPoints := make([]string, 0, len(mounts.Mounts))
	for _, m := range mounts.Mounts {
		if validFS.MatchString(m.FSType) {
			mountPoints = append(mountPoints, strings.TrimSpace(m.MountPoint))
		}
	}
	return &DiskLinuxResource{mountPoints}
}

func (r DiskLinuxResource) Init() *InitPlotMesage {
	return &InitPlotMesage{Id: "dk", Type: "init", Label: r.MountPoints, Min: 0.0, Max: 100.0}
}

func (r DiskLinuxResource) Probe() (*UpdateMessage, error) {
	values := make([]float64, 0, len(r.MountPoints))
	for _, m := range r.MountPoints {
		disk, err := linuxproc.ReadDisk(m)
		if err != nil {
			panic("Cannot read stat for " + m)
		}
		used_percent := (float64(disk.Used) * 100.0) / float64(disk.All)
		values = append(values, used_percent)
	}
	return &UpdateMessage{Id: "dk", V: values}, nil
}
