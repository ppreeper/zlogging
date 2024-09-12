package main

import (
	"io"
	"log/slog"
	"testing"

	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BenchmarkSLog(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(
			"incoming request",
			"method", "GET",
			"time_taken_ms", 158,
			"path", "/hello/world?q=search",
			"status", 200,
			"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
		)
	}
}

func BenchmarkZAP(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zap.InfoLevel,
	)
	logger := zap.New(core)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("incoming request",
			zap.String("method", "GET"),
			zap.Int("time_taken_ms", 158),
			zap.String("provider", "google"),
			zap.String("path", "/hello/world?q=search"),
			zap.Int("status", 200),
			zap.String("user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)"),
		)
	}
}

func BenchmarkZAPSugar(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zap.InfoLevel,
	)
	logger := zap.New(core)
	sugar := logger.Sugar()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sugar.Infow("incoming request",
			"method", "GET",
			"time_taken_ms", 158,
			"path", "/hello/world?q=search",
			"status", 200,
			"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
		)
	}
}

func BenchmarkZerolog(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := zerolog.New(io.Discard).With().Timestamp().Logger()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().
			Str("method", "GET").
			Int("time_taken_ms", 158).
			Str("path", "/hello/world?q=search").
			Int("status", 200).
			Str("user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)").
			Msg("incoming request")
	}
}
