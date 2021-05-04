package log

import "github.com/utilbox/gcommon/gson"

type LogConfig struct {
	Level      string
	Depth      int
	Path       string
	MaxSize    int // in MegaBytes
	MaxBackups int
	MaxAge     int // in Days
	Compress   bool
	LocalTime  bool
}

func (l *LogConfig) SetDefault() {
	if l.Level == "" {
		l.Level = "Debug"
	}
	if l.Depth == 0 {
		l.Depth = 1
	}
	if l.Path == "" {
		l.Path = "./log/log"
	}
	if l.MaxSize == 0 {
		l.MaxSize = 10
	}
	if l.MaxBackups == 0 {
		l.MaxBackups = 100
	}
	if l.MaxAge == 0 {
		l.MaxAge = 7
	}
}

func (l *LogConfig) String() string {
	if l == nil {
		return "nil"
	}

	raw, _ := gson.Marshal(l)
	return string(raw)
}
