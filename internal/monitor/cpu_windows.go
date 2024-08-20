// cpu_windows.go
//go:build windows

package monitor

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/kafka-embracetheday/goResourceWatcher/internal/logger"
)

var log = logger.GetLogger()

type CPUUsage struct {
}

func (c *CPUUsage) getCPUUsage() (float64, error) {
	idle1, kernel1, user1, err := c.getSystemTimes()
	if err != nil {
		log.Errorf("get windows system time error:%s", err)
		return 0, err
	}

	time.Sleep(1000 * time.Millisecond)

	idle2, kernel2, user2, err := c.getSystemTimes()
	if err != nil {
		log.Errorf("get windows system time error:%s", err)
		return 0, err
	}

	totalIdle := idle2 - idle1
	totalKernel := kernel2 - kernel1
	totalUser := user2 - user1

	total := totalKernel + totalUser
	used := total - totalIdle

	return float64(used) / float64(total) * 100, nil
}

func (c *CPUUsage) getSystemTimes() (idle, kernel, user uint64, err error) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")           //kernel32.dll包是windows操作系统的一个库
	procGetSystemTimes := kernel32.NewProc("GetSystemTimes") // GetSystemTimes获取cpu时间

	var idleTime, kernelTime, userTime syscall.Filetime

	ret, _, err := procGetSystemTimes.Call(
		uintptr(unsafe.Pointer(&idleTime)),
		uintptr(unsafe.Pointer(&kernelTime)),
		uintptr(unsafe.Pointer(&userTime)),
	)
	if ret == 0 {
		log.Infof("don't get windows system times")
		return 0, 0, 0, err
	}

	idle = c.fileTimeToUint64(&idleTime)
	kernel = c.fileTimeToUint64(&kernelTime)
	user = c.fileTimeToUint64(&userTime)

	return idle, kernel, user, nil
}

func (c *CPUUsage) fileTimeToUint64(ft *syscall.Filetime) uint64 {
	return uint64(ft.HighDateTime)<<32 + uint64(ft.LowDateTime)
}
