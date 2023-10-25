package linux

import "github.com/shirou/gopsutil/v3/disk"

type LogicalDisk struct {
	Name           string
	Size           int
	FreeSpace      int
	PercentageFree float32
}

func Disk() (LogicalDisk, error) {
	var d LogicalDisk

	rootdisk, err := disk.Usage("/")
	if err != nil {
		return LogicalDisk{}, err
	}

	d.Name = rootdisk.Path
	d.Size = int(rootdisk.Total)
	d.FreeSpace = int(rootdisk.Free)
	d.PercentageFree = float32(rootdisk.UsedPercent)

	return d, nil
}
