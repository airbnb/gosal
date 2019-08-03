package common

import "github.com/shirou/gopsutil/mem"

func GetOS() (OS, error) {
	v, _ := mem.VirtualMemory()

	memory := OS{
		Caption:                "undefined",
		TotalVirtualMemorySize: v.Total,
		TotalVisibleMemorySize: v.Total / 1024, // Need to debug.
	}

	return memory, nil
}

type OS struct {
	Caption                string
	TotalVirtualMemorySize uint64
	TotalVisibleMemorySize uint64
}
