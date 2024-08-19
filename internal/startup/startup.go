package startup

import (
	"github.com/kafka-embracetheday/goResourceWatcher/config"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/logger"
	"os"
)

type Server struct {
}

func (s *Server) StartUp() {
	// init config
	config.LoadConfig()
	// init logger
	logger.InitLogger()

	log := logger.GetLogger()
	env := os.Getenv("APP_ENV")
	log.Infof("Project is running, the project environment is %s", env)
}
