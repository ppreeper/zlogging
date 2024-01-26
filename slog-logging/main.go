package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
)

// https://betterstack.com/community/guides/logging/logging-in-go/

func main() {
	logs := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logs)
	// slog.Debug("Debug message")
	// slog.Info("Info message")
	// slog.Warn("Warning message")
	// slog.Error("Error message")
	defaultLogger()

	contextualAttributes()

	childLoggers()

	loggerLevels()

	loggerLevelsPostlevel()

	colorLog()

	contextWithSlog()
	contextWithSlog2()

	hidingSensitiveData()
}

func defaultLogger() {
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")
}

func contextualAttributes() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info(
		"incoming request",
		"method", "GET",
		"time_taken_ms", 158,
		"path", "/hello/world?q=search",
		"status", 200,
		"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
	)
	logger.Info(
		"incoming request",
		slog.String("method", "GET"),
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		slog.Int("status", 200),
		slog.String(
			"user_agent",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	)
	logger.Info(
		"incoming request",
		"method", "GET",
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		"status", 200,
		slog.String(
			"user_agent",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	)
	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"image uploaded",
		slog.Int("id", 23123),
		slog.Group("properties",
			slog.Int("width", 4000),
			slog.Int("height", 3000),
			slog.String("format", "jpeg"),
		),
	)
}

func childLoggers() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	buildInfo, _ := debug.ReadBuildInfo()

	logger := slog.New(handler)

	child := logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
		),
	)
	logger.Info("logger image upload successful", slog.String("image_id", "39ud88"))
	child.Info("child image upload successful", slog.String("image_id", "39ud88"))
	child.Warn(
		"storage is 90% full",
		slog.String("available_space", "900.1 mb"),
	)
}

func loggerLevels() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")
}

func loggerLevelsPostlevel() {
	logLevel := &slog.LevelVar{} // INFO

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	logLevel.Set(slog.LevelDebug)

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")
}

func contextWithSlog() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx := context.WithValue(context.Background(), "request_id", "req-123")

	logger.InfoContext(ctx, "image uploaded", slog.String("image_id", "img-998"))
}

func contextWithSlog2() {
	h := &ContextHandler{slog.NewJSONHandler(os.Stdout, nil)}

	logger := slog.New(h)

	ctx := AppendCtx(context.Background(), slog.String("request_id", "req-123"))

	logger.InfoContext(ctx, "image uploaded", slog.String("image_id", "img-998"))
}

func colorLog() {
	opts := PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := NewPrettyHandler(os.Stdout, opts)
	logger := slog.New(handler)
	logger.Debug(
		"executing database query",
		slog.String("query", "SELECT * FROM users"),
	)
	logger.Info("image upload successful", slog.String("image_id", "39ud88"))
	logger.Warn(
		"storage is 90% full",
		slog.String("available_space", "900.1 MB"),
	)
	logger.Error(
		"An error occurred while processing the request",
		slog.String("url", "https://example.com"),
	)
}

func hidingSensitiveData() {
	fmt.Println("Hiding sensitive data")

	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)

	u := &User{
		ID:        "user-12234",
		FirstName: "Jan",
		LastName:  "Doe",
		Email:     "jan@example.com",
		Password:  "pass-12334",
	}

	logger.Info("info", "user", u)
}

// User does not implement `LogValuer` here
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// implement the `LogValuer` interface
func (u *User) LogValue() slog.Value {
	return slog.StringValue(u.ID)
}
