package log

import (
	"io"
	"io/ioutil"
	"os"
)

const (
	OFF = iota
	DEBUG
	INFO
	WARNING
	ERROR
)

const (
	MinLevel = OFF
	MaxLevel = ERROR
)

func GetPrefixForLogger(level int) string {
	switch level {
	case DEBUG:
		return "[DEBUG] "
	case INFO:
		return "[INFO] "
	case WARNING:
		return "[WARNING] "
	case ERROR:
		return "[ERROR] "
	default:
		return ""
	}
}

func GetValueFromParam(param string) int {
	switch param {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARNING":
		return WARNING
	case "ERROR":
		return ERROR
	default:
		return OFF
	}
}

func GetIOHandle(level int) io.Writer {
	switch level {
	case DEBUG:
		return os.Stderr
	case INFO:
		return os.Stderr
	case WARNING:
		return os.Stderr
	case ERROR:
		return os.Stderr
	default:
		return ioutil.Discard
	}
}
