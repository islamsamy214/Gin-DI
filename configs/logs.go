package configs

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
)

func NewLogsWriterConfig() io.Writer {
	appConfig := NewAppConfig()
	dailyLogger := &lumberjack.Logger{
		Filename:   "./storage/logs/" + appConfig.Name + "-" + time.Now().Format("2006-01-02") + ".log",
		MaxSize:    10, // Max megabytes before log is rotated
		MaxBackups: 7,  // Max number of old log files to retain
		MaxAge:     30, // Max number of days to retain log files
		Compress:   true,
	}

	// Combine daily log file and console for output
	multiWriter := io.MultiWriter(os.Stdout, dailyLogger)

	// Create a standard logger
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return multiWriter
}
