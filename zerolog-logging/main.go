package main

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
)

// https://betterstack.com/community/guides/logging/zerolog/

func main() {
	// zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Info().Msg("info message")
	logger.Debug().Str("username", "joshua").Send()
	logger.Info().Msg("Info message")
	logger.Error().Msg("Error message")
	logger.Info().
		Str("name", "john").
		Int("age", 22).
		Bool("registered", true).
		Msg("new signup!")
	logger.Info().
		Str("name", "john").
		Int("age", 22).
		Bool("registered", true).
		Send()

	buildInfo, _ := debug.ReadBuildInfo()

	logger2 := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()

	logger2.Trace().Msg("trace message")
	logger2.Debug().Msg("debug message")
	logger2.Info().Msg("info message")
	logger2.Warn().Msg("warn message")
	logger2.Error().Msg("error message")
	logger2.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	logger2.WithLevel(zerolog.PanicLevel).Msg("panic message")

	logger3 := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger().
		Sample(&zerolog.BasicSampler{N: 5})

	for i := 1; i <= 10; i++ {
		logger3.Info().Msgf("a message from the gods: %d", i)
	}

	l := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger().
		Sample(&zerolog.BurstSampler{
			Burst:  3,
			Period: 1 * time.Second,
		})

	for i := 1; i <= 10; i++ {
		l.Info().Msgf("a message from the gods: %d", i)
		l.Warn().Msgf("warn message: %d", i)
		l.Error().Msgf("error message: %d", i)
	}

	infoSampler := &zerolog.BurstSampler{
		Burst:  3,
		Period: 1 * time.Second,
	}

	warnSampler := &zerolog.BurstSampler{
		Burst:  3,
		Period: 1 * time.Second,
		// Log every 5th message after exceeding the burst rate of 3 messages per
		// second
		NextSampler: &zerolog.BasicSampler{N: 5},
	}

	errorSampler := &zerolog.BasicSampler{N: 2}

	l2 := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger().
		Sample(zerolog.LevelSampler{
			WarnSampler:  warnSampler,
			InfoSampler:  infoSampler,
			ErrorSampler: errorSampler,
		})

	for i := 1; i <= 10; i++ {
		l2.Info().Msgf("a message from the gods: %d", i)
		l2.Warn().Msgf("warn message: %d", i)
		l2.Error().Msgf("error message: %d", i)
	}

	file, err := os.OpenFile(
		"myapp.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o664,
	)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	loggerFile := zerolog.New(file).With().Timestamp().Logger()

	loggerFile.Info().Msg("Info message")
}
