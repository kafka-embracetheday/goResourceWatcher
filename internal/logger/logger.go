package logger

import (
	"github.com/kafka-embracetheday/goResourceWatcher/config"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"

	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var Logger *logrus.Logger

func GetLogger() *logrus.Logger { return Logger }

type LogEntry struct {
	Time  string `json:"time"`
	Level string `json:"level"`
	Msg   string `json:"msg"`
	File  string `json:"file"`
	Line  int    `json:"line"`
	Func  string `json:"func"`
}

func InitLogger() {

	logConfig := config.GetConfig().Logger

	Logger = logrus.New()

	level, err := logrus.ParseLevel(logConfig.Level)
	if err != nil {
		fmt.Printf("Invalid log level: %s, defaulting to InfoLevel\n", level)
		level = logrus.InfoLevel
	}

	Logger.SetLevel(level)
	Logger.SetFormatter(&CustomJSONFormatter{
		TimestampFormat: time.RFC3339,
		EnableColors:    true,
	})

	// 创建 log 目录如果不存在
	logDir := logConfig.Path
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating log directory: %v\n", err)
		return
	}

	// 生成带时间戳的日志文件名
	timestamp := time.Now().Format("20060102") // 格式化时间戳
	logFileName := filepath.Join(logDir, fmt.Sprintf("log-%s.log", timestamp))

	Logger.SetOutput(&lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    50,   // MB
		MaxBackups: 3,    // 保留旧文件的最大个数
		MaxAge:     30,   // 天
		Compress:   true, // 是否压缩
	})
}

type CustomJSONFormatter struct {
	TimestampFormat string
	EnableColors    bool
}

func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取调用者信息
	_, file, line, ok := runtime.Caller(7)
	if ok {
		file = filepath.Base(file)
	}

	level := entry.Level.String()
	if f.EnableColors {
		level = getColoredLevel(level)
	}

	logEntry := LogEntry{
		Time:  entry.Time.Format(f.TimestampFormat),
		Level: entry.Level.String(),
		Msg:   entry.Message,
		File:  file,
		Line:  line,
		Func:  getFunctionName(entry),
	}

	data, err := json.Marshal(logEntry)
	if err != nil {
		return nil, err
	}

	return append(data, '\n'), nil
}

// 获取函数名
func getFunctionName(entry *logrus.Entry) string {
	if pc, _, _, ok := runtime.Caller(8); ok {
		return runtime.FuncForPC(pc).Name()
	}
	return ""
}

func getColoredLevel(level string) string {
	switch level {
	case "debug":
		return "\033[34m" + level + "\033[0m" // 蓝色
	case "info":
		return "\033[32m" + level + "\033[0m" // 绿色
	case "warn":
		return "\033[33m" + level + "\033[0m" // 黄色
	case "error":
		return "\033[31m" + level + "\033[0m" // 红色
	case "fatal":
		return "\033[35m" + level + "\033[0m" // 紫色
	case "panic":
		return "\033[41m" + level + "\033[0m" // 背景红色
	default:
		return level
	}
}
