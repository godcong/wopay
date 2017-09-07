package wxpay

import (
	"io"
	"log"
	"os"
	"strings"
	"time"

	"runtime"
)

const LOG_PATH = "log"
const SYSTEM_SEPARATOR_WINDOWS = "\\"
const SYSTEM_SEPARATOR_LINUX = "/"

func init() {
	logConfig()
}

func GetLogPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return LOG_PATH
	}
	return strings.Join([]string{dir, LOG_PATH}, getSystemSeparator())
}

func logConfig() {
	log.Println("log init")
	name := strings.Join([]string{"log", time.Now().Format("20060102")}, "_")
	name = strings.Join([]string{GetLogPath(), name}, getSystemSeparator()) + ".log"
	os.MkdirAll(GetLogPath(), os.ModePerm)
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, os.ModePerm)
	if err != nil {
		return
	}

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func getSystemSeparator() string {
	if runtime.GOOS == "linux" {
		return SYSTEM_SEPARATOR_LINUX
	}
	return SYSTEM_SEPARATOR_WINDOWS
}
