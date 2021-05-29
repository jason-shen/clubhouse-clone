package utils

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var log zerolog.Logger

type Level uint8

const DebugLevel Level = iota

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log = zerolog.New(output).With().Timestamp().Logger()
	SetLevel(DebugLevel)
}

func SetLevel(l Level)  {
	zerolog.SetGlobalLevel(zerolog.Level(l))
}

func Infof(format string, v ...interface{})  {
	log.Info().Msgf(format, v...)
}

func Debugf(format string, v ...interface{})  {
	log.Debug().Msgf(format, v...)
}

func Warnf(format string, v ...interface{})  {
	log.Warn().Msgf(format, v...)
}

func Errorf(format string, v ...interface{})  {
	log.Error().Msgf(format, v...)
}

func Panicf(format string, v ...interface{})  {
	log.Panic().Msgf(format, v...)
}

func Fatalf(format string, v ...interface{})  {
	log.Fatal().Msgf(format, v...)
}

