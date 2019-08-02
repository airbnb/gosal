package common

import (
	"github.com/shirou/gopsutil/cpu"
)

func GetProcessor() (Processor, error) {
	c, _ := cpu.Info()

	cpu := Processor{
		CPUType:               c[0].ModelName,
		CurrentProcessorSpeed: int(c[0].Mhz),
	}

	return cpu, nil
}

type Processor struct {
	CPUType               string
	CurrentProcessorSpeed int
}
