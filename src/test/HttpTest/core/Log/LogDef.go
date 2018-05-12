package Log

import "strings"

const (
	LOG_ANY = 0
	LOG_TRACE = 1
	LOG_DEBUG = 2
	LOG_ERROR = 3
	LOG_FATAL = 4
	LOG_NONE = 100
)

const (
	ShiftFileSize = 1024 *1024*40
	ShiftFileCount = 10
	CloseTimeoutLogFileTime = 300
)

func LevelToString(logLevel int) string {
	switch logLevel {
	case LOG_TRACE:
		return "T"
	case LOG_DEBUG:
		return "D"
	case LOG_ERROR:
		return "E"
	case LOG_FATAL:
		return "F"
	}
	return ""
}

func ParseFromString(strLevel string) int {
	if len(strLevel)== 0 { return LOG_DEBUG}
	var lowStrLevel string = strings.ToLower(strLevel)
	if lowStrLevel == "trace" {
		return LOG_TRACE
	}else if lowStrLevel == "debug" {
		return LOG_DEBUG
	}else if lowStrLevel == "error" {
		return LOG_ERROR
	}else if lowStrLevel == "fatal" {
		return  LOG_FATAL
	}else if lowStrLevel == "none" {
		return LOG_NONE
	}
	return LOG_DEBUG
}

