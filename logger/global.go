package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

const (
	NormalFormat = "normal"
	JsonFormat   = "json"
)

// InitLogger initiates the global logger
func InitLogger(logLevel string, logFormat string, filePath string, enableConsole bool, noColor bool) error {
	logFile := io.Discard
	consoleOutput := io.Discard

	var consoleWriter io.Writer
	var fileWriter io.Writer

	if filePath != "" {
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		logFile = f
	}

	if enableConsole {
		consoleOutput = os.Stdout
	}

	if err := setLogLevel(logLevel); err != nil {
		return err
	}

	if logFormat == NormalFormat {
		normalFileWriter.Out = logFile
		normalConsoleWriter.Out = consoleOutput

		if noColor {
			normalConsoleWriter.NoColor = true
		}

		consoleWriter = normalConsoleWriter
		fileWriter = normalFileWriter

	} else if logFormat == JsonFormat {
		fileWriter = logFile
		consoleWriter = consoleOutput
	} else {
		return fmt.Errorf("log format '%s' is not supported (json, normal)", logFormat)
	}
	logger = zerolog.New(zerolog.MultiLevelWriter(consoleWriter, fileWriter)).With().Timestamp().Logger()
	return nil
}

func Debug(msg string) {
	logger.Debug().Msg(msg)
}
func Debugf(msg string, v ...interface{}) {
	logger.Debug().Msgf(msg, v...)
}

func Info(msg string) {
	logger.Info().Msg(msg)
}
func Infof(msg string, v ...interface{}) {
	logger.Info().Msgf(msg, v...)
}

func Error(err error, msg string) error {
	if err != nil {
		logger.Error().Str("error", err.Error()).Msg(msg)
	} else {
		logger.Error().Msgf(msg)
	}
	return err
}
func Errorf(err error, msg string, v ...interface{}) error {
	if err != nil {
		logger.Error().Str("error", err.Error()).Msgf(msg, v...)
	} else {
		logger.Error().Msgf(msg, v...)
	}
	return err
}

func Warn(msg string) {
	logger.Warn().Msg(msg)
}
func Warnf(msg string, v ...interface{}) {
	logger.Warn().Msgf(msg, v...)
}

func Panic(msg string) {
	logger.Panic().Msg(msg)
}
func Panicf(msg string, v ...interface{}) {
	logger.Panic().Msgf(msg, v...)
}
