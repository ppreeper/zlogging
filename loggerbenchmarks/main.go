package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	fmt.Println("Logger Benchmark")

	// "incoming request",
	// "method", "GET",
	// "time_taken_ms", 158,
	// "path", "/hello/world?q=search",
	// "status", 200,
	// "user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",

	// output := io.Discard
	output := os.Stdout

	slogger := slog.New(slog.NewJSONHandler(output, nil))

	cfg := zap.NewProductionConfig()
	zaplogger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(output),
		zap.InfoLevel,
	))

	sugarlogger := zaplogger.Sugar()

	zerologger := zerolog.New(output).With().Timestamp().Logger()

	fmt.Println("slog")
	slogger.Info(
		"incoming request",
		"method", "GET",
		"time_taken_ms", 158,
		"path", "/hello/world?q=search",
		"status", 200,
		"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
	)

	fmt.Println("zap")
	zaplogger.Info("incoming request",
		zap.String("method", "GET"),
		zap.Int("time_taken_ms", 158),
		zap.String("provider", "google"),
		zap.String("path", "/hello/world?q=search"),
		zap.Int("status", 200),
		zap.String("user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)"),
	)
	fmt.Println("zap sugar")
	sugarlogger.Infow("incoming request",
		"method", "GET",
		"time_taken_ms", 158,
		"path", "/hello/world?q=search",
		"status", 200,
		"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
	)
	fmt.Println("zerolog")
	zerologger.Info().
		Str("method", "GET").
		Int("time_taken_ms", 158).
		Str("path", "/hello/world?q=search").
		Int("status", 200).
		Str("user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)").
		Msg("incoming request")
}
