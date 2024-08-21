package main

import (
	"context"
	"fmt"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/monitor"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/startup"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/task"
	"sync"
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
	server.HandleSignal()
	fmt.Print(watcher)

	// task使用方法
	tq := task.NewTaskQueue()

	taskId := tq.AddTask("running", myTask)

	fmt.Println(taskId)

	time.Sleep(3 * time.Second)
	tq.PauseTask(taskId)

	time.Sleep(10 * time.Second)
	tq.ResumeTask(taskId)

	select {}
}

// task例子
func myTask(ctx context.Context, pause *sync.Cond, isPaused *bool) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("任务结束")
			return
		default:
			pause.L.Lock()
			for *isPaused {
				pause.Wait() // 等待被唤醒
			}
			pause.L.Unlock()

			// 执行任务逻辑
			cpuMonitor := monitor.NewCPUMonitor()
			usage, err := cpuMonitor.GetCPUUsage()
			if err != nil {
				fmt.Println("get cpu usage error:", err)
				return
			}
			fmt.Println("cpu usage:", usage)
			time.Sleep(100 * time.Millisecond) // 模拟任务执行
		}
	}
}
