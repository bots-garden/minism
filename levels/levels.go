package levels

import "github.com/extism/extism"

func GetLevel(logLevel string) extism.LogLevel {
	level := extism.Off
	switch logLevel {
	case "error":
		level = extism.Error
	case "warn":
		level = extism.Warn
	case "info":
		level = extism.Info
	case "debug":
		level = extism.Debug
	case "trace":
		level = extism.Trace
	}
	return level
}
