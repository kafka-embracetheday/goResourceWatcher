package startup

import (
	"fmt"
	"github.com/kafka-embracetheday/goResourceWatcher/config"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/logger"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/mysql"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/task"
	"os/signal"
	"syscall"

	"os"
	"runtime"
)

type Server struct {
}

func (s *Server) StartUp() {
	// init config
	config.LoadConfig()
	// init logger
	logger.InitLogger()
	// init mysql
	mysql.InitMysql()
	// init task entry
	task.AutoMigrateTaskEntity()

	env := os.Getenv("APP_ENV")
	goos := runtime.GOOS
	logger.Logger.Infof("Project is running, the project environment is %s, the project os is %s", env, goos)
}

func (s *Server) Shutdown() {
	err := task.ClearTaskTable()
	if err != nil {
		logger.Logger.Error("Failed to task clear task table:", err)
		return
	}
}

func (s *Server) HandleSignal() {
	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGHUP)
	go func() {
		fmt.Println("接受关闭信号")
		<-closeChan
		s.Shutdown()
		os.Exit(0)
	}()
}
