package main

import (
	"fmt"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/logger"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/monitor"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/startup"
	"time"
)

const watcher = `
   __
  /  \
 |    |
  \__/
Watcher is active
`

func main() {
	server := startup.Server{}
	server.StartUp()

	fmt.Print(watcher)
	log := logger.GetLogger()

	for {
		cpuMonitor := monitor.NewCPUMonitor()
		usage, err := cpuMonitor.GetCPUUsage()
		if err != nil {
			log.Errorf("获取 CPU 占用率失败:%s", err)
			return
		}

		fmt.Printf("CPU 占用率: %.2f%%\n", usage)
		time.Sleep(1 * time.Second)
	}
}
