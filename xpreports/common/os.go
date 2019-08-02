package common

import "github.com/shirou/gopsutil/mem"

func GetOS() (OS, error) {

	v, _ := mem.VirtualMemory()

	memory := OS{
		Caption:                "undefined",
		TotalVirtualMemorySize: int(v.Total),
		TotalVisibleMemorySize: int(v.SwapTotal) + int(v.Total),
	}

	return memory, nil

}

type OS struct {
	Caption                string
	TotalVirtualMemorySize int
	TotalVisibleMemorySize int
}
