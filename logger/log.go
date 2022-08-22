package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

type (
	LoggerStruct struct {
		*logrus.Logger
	}
)

var logger *LoggerStruct

func StartLogger() *LoggerStruct {
	var baseLogger = logrus.New()
	loggerStruct := &LoggerStruct{baseLogger}
	textFormattor := &logrus.TextFormatter{}
	textFormattor.DisableColors = false
	textFormattor.TimestampFormat = time.ANSIC
	loggerStruct.Formatter = textFormattor
	logger = loggerStruct
	return loggerStruct
}
func GetLogger() *LoggerStruct {
	if logger == nil {
		return StartLogger()
	}
	return logger
}

func (logger *LoggerStruct) InfoLog(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func (logger *LoggerStruct) PrintLog(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

func (logger *LoggerStruct) WarnLog(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func (logger *LoggerStruct) WarningLog(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}
func (logger *LoggerStruct) WarningLn(args ...interface{}) {
	logger.Warningln(args...)
}
func (logger *LoggerStruct) ErrorLog(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func (logger *LoggerStruct) FatalLog(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func (logger *LoggerStruct) PanicLog(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
