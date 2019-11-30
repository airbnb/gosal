package linux

import "github.com/shirou/gopsutil/disk"

type LogicalDisk struct {
	Name      string
	Size      int
	FreeSpace int
}

func GetDisk() (LogicalDisk, error) {
	var d LogicalDisk

	rootdisk, _ := disk.Usage("/")

	d.Name = rootdisk.Path
	d.Size = int(rootdisk.Total)
	d.FreeSpace = int(rootdisk.Free)

	return d, nil
}
