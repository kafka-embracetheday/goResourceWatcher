package monitor

type CPUMonitorStrategy interface {
	getCPUUsage() (float64, error)
}

type CPUMonitor struct {
	strategy CPUMonitorStrategy
}

func NewCPUMonitor() *CPUMonitor {
	var strategy CPUMonitorStrategy
	strategy = &CPUUsage{}
	return &CPUMonitor{strategy: strategy}
}

func (c *CPUMonitor) GetCPUUsage() (float64, error) {
	return c.strategy.getCPUUsage()
}
