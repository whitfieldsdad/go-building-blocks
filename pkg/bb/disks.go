package bb

import (
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/disk"
)

type Disk struct {
	Device     string          `json:"device"`
	Usage      *DiskUsage      `json:"disk_usage,omitempty"`
	Partitions []DiskPartition `json:"partitions,omitempty"`
}

func (d Disk) GetDiskUsage() (*DiskUsage, error) {
	total := uint64(0)
	used := uint64(0)
	free := uint64(0)
	for _, partition := range d.Partitions {
		if partition.DiskUsage == nil {
			continue
		}
		total += partition.DiskUsage.Total
		used += partition.DiskUsage.Used
		free += partition.DiskUsage.Free
	}
	if total == 0 {
		return nil, errors.New("failed to get disk usage")
	}
	return &DiskUsage{
		Total:       total,
		Used:        used,
		Free:        free,
		UsedPercent: float64(used) / float64(total),
	}, nil
}

type DiskPartition struct {
	Device         string     `json:"device"`
	Mountpoint     string     `json:"mountpoint"`
	FilesystemType string     `json:"filesystem_type"`
	Options        []string   `json:"options"`
	DiskUsage      *DiskUsage `json:"disk_usage"`
}

type DiskUsage struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

// List disks returns a list of mounted disks.
func ListDisks() ([]Disk, error) {
	devices, err := disk.Partitions(true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list disk partitions")
	}
	disks := make([]Disk, 0)
	for _, device := range devices {
		if device.Device == "none" {
			continue
		}
		disk, err := GetDisk(device.Mountpoint)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get disk")
		}
		disks = append(disks, *disk)
	}
	return disks, nil
}

// ListDiskPartitions returns a list of mounted disk partitions.
func ListDiskPartitions() ([]DiskPartition, error) {
	rows, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}
	devices := make(map[string]bool)
	partitions := make([]DiskPartition, len(rows))
	for i, row := range rows {
		partitions[i] = DiskPartition{
			Device:         row.Device,
			Mountpoint:     row.Mountpoint,
			FilesystemType: row.Fstype,
			Options:        row.Opts,
		}
		devices[row.Device] = true
	}
	return partitions, nil
}

// GetDisk returns information about a disk.
func GetDisk(path string) (*Disk, error) {
	usage, err := GetDiskUsage(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get disk usage")
	}
	return &Disk{
		Device: path,
		Usage:  usage,
	}, nil
}

// GetDiskUsage returns information about disk usage.
func GetDiskUsage(path string) (*DiskUsage, error) {
	stat, err := disk.Usage(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get disk usage")
	}
	return &DiskUsage{
		Total:       stat.Total,
		Used:        stat.Used,
		Free:        stat.Free,
		UsedPercent: stat.UsedPercent,
	}, nil
}
