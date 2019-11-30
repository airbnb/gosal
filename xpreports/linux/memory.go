package linux

import (
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func GetOS() (OS, error) {
	v, _ := mem.VirtualMemory()
	h, _ := host.Info()

	memory := OS{
		Caption:                h.PlatformVersion,
		TotalVirtualMemorySize: v.Total,
		TotalVisibleMemorySize: v.Total / 1024,
	}
	return memory, nil
}

type OS struct {
	Caption                string
	TotalVirtualMemorySize uint64
	TotalVisibleMemorySize uint64
}
