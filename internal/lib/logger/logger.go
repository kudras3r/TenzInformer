package logger

import (
	"log"
	"os"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct {
	level  int
	logger *log.Logger
}

func NewLogger(level int, logFile string) (*Logger, error) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	return &Logger{
		level:  level,
		logger: log.New(file, "", log.Ldate|log.Ltime),
	}, nil
}

func (l *Logger) write(level int, message string) {
	var prefix string

	if level >= l.level {
		switch level {
		case DEBUG:
			prefix = "< DEBUG > "
		case INFO:
			prefix = "< INFO > "
		case WARN:
			prefix = "< WARN > "
		case ERROR:
			prefix = "< ERROR > "
		case FATAL:
			prefix = "< FATAL > "
		}

		l.logger.Println(prefix + message)
		if level == FATAL {
			os.Exit(0)
		}
	}
}

func (l *Logger) DEBUG(message string) {
	l.write(DEBUG, message)
}

func (l *Logger) INFO(message string) {
	l.write(INFO, message)
}

func (l *Logger) WARN(message string) {
	l.write(WARN, message)
}

func (l *Logger) ERROR(message string) {
	l.write(ERROR, message)
}

func (l *Logger) FATAL(message string) {
	l.write(FATAL, message)
}
