package Metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)

var (
	CPUPercentage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cpu_percentage",
			Help: "Current CPU usage percentage",
		},
	)

	MemoryUsed = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "memory_used",
			Help: "Current memory usage in bytes",
		},
	)

	DiskUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "disk_usage",
			Help: "Current disk usage in bytes",
		},
		[]string{"disk"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(CPUPercentage)
	prometheus.MustRegister(MemoryUsed)
	prometheus.MustRegister(DiskUsage)

	go func() {
		for {
			updateMetrics()
			time.Sleep(1 * time.Second)
		}
	}()
}

func updateMetrics() {
	cpuPercent, err := cpu.Percent(time.Second, true)
	if err == nil && len(cpuPercent) > 0 {
		CPUPercentage.Set(cpuPercent[0])
	}

	memInfo, err := mem.VirtualMemory()
	if err == nil {
		MemoryUsed.Set(float64(memInfo.Used))
	}

	partitions, err := disk.Partitions(false)
	if err != nil {
		return
	}

	for _, partition := range partitions {
		usageStat, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}
		DiskUsage.WithLabelValues(partition.Device).Set(float64(usageStat.Used))
	}
}
