package wtools

import (
	"fmt"
	"github.com/charmbracelet/log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	logger     *log.Logger
	fileLogger *log.Logger
)

func LogToFile() error {
	var (
		logDir  = "/var/log/vda/xlog"
		logFile = "vda-rollup.log"
	)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		fmt.Println("日志目录不存在，使用当前目录下的 vda-rollup.log")
		logFile = "./" + logFile
	} else {
		fmt.Println("日志目录存在，日志将输出到:", logFile)
		logFile = filepath.Join(logDir, logFile)
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("无法打开日志文件: %w", err)
	}

	fileLogger = log.NewWithOptions(file, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.RFC3339,
		Prefix:          "<Will Yin 🚀>",
	})
	fileLogger.SetFormatter(log.TextFormatter)

	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.RFC3339,
		Prefix:          "<Will Yin 🚀>",
	})
	logger.SetFormatter(log.TextFormatter)

	return nil
}

func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown:0"
	}
	shortFile := filepath.Base(file)
	return fmt.Sprintf("%s:%d", shortFile, line)
}

func Info(msg string) {
	caller := getCallerInfo()
	logger.Info(msg, "caller", caller)
	fileLogger.Info(msg, "caller", caller)
}

func Error(msg string) {
	caller := getCallerInfo()
	logger.Error(msg, "caller", caller)
	fileLogger.Error(msg, "caller", caller)
}

func Debug(msg string) {
	caller := getCallerInfo()
	logger.Debug(msg, "caller", caller)
	fileLogger.Debug(msg, "caller", caller)
}

func Warn(msg string) {
	caller := getCallerInfo()
	logger.Warn(msg, "caller", caller)
	fileLogger.Warn(msg, "caller", caller)
}
