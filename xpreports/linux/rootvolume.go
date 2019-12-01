package linux

import "github.com/shirou/gopsutil/disk"

type LogicalDisk struct {
	Name           string
	Size           int
	FreeSpace      int
	PercentageFree float32
}

func GetDisk() (LogicalDisk, error) {
	var d LogicalDisk

	rootdisk, _ := disk.Usage("/")

	d.Name = rootdisk.Path
	d.Size = int(rootdisk.Total)
	d.FreeSpace = int(rootdisk.Free)
	d.PercentageFree = float32(rootdisk.UsedPercent)

	return d, nil
}
