package logger

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	level := entry.Level.String()
	message := entry.Message

	level = fmt.Sprintf("[%s]", stringUpper(level))

	logLine := fmt.Sprintf("%s %s - %s\n", timestamp, level, message)
	return []byte(logLine), nil
}

func Init(appName string) {
	logFilePath, err := getLogFilePath(appName)
	if err != nil {
		fmt.Printf("Error configuring logs: %v\n", err)
		os.Exit(1)
	}

	file, _ := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	log.SetOutput(file)
	log.SetFormatter(&CustomFormatter{})
	log.SetLevel(log.InfoLevel)
}

func getLogFilePath(appName string) (string, error) {
	var logDir string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Couldn't open the user home directory.")
	}

	systemLogDir := filepath.Join("/var/log", appName)
	userLogDir := filepath.Join(
		homeDir,
		".local",
		"share",
		appName,
		"logs",
	)

	if err := os.MkdirAll(systemLogDir, 0755); err == nil {
		testFile := filepath.Join(systemLogDir, ".testwrite")
		if err := os.WriteFile(testFile, []byte("ok"), 0644); err == nil {
			_ = os.Remove(testFile)
			logDir = systemLogDir
		} else {
			logDir = userLogDir
		}
	} else {
		logDir = userLogDir
	}

	if err := os.MkdirAll(logDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create logs directory: %w", err)
	}

	logFile := fmt.Sprintf("%s.log", appName)

	return filepath.Join(logDir, logFile), nil
}

func stringUpper(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c = c - 32
		}
		result += string(c)
	}
	return result
}
