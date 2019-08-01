package darwin

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
	d.Size = int(rootdisk.Total / 1024)
	d.FreeSpace = int(rootdisk.Free / 1024)

	// print(strconv.Itoa(d.Size))
	return d, nil
}
