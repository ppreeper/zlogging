package main

import (
	"time"

	"go.uber.org/zap"
)

// https://betterstack.com/community/guides/logging/go/zap/
func main() {
	logger := zap.Must(zap.NewProduction())

	defer logger.Sync()

	sugar := logger.Sugar()

	logger.Info("Hello from Zap logger!")

	logger.Info("User logged in",
		zap.String("username", "johndoe"),
		zap.Int("userid", 123456),
		zap.String("provider", "google"),
	)

	sugar.Info("Hello from Zap logger!")
	sugar.Infoln(
		"Hello from Zap logger!",
	)
	sugar.Infof(
		"Hello from Zap logger! The time is %s",
		time.Now().Format("03:04 AM"),
	)

	sugar.Infow("User logged in",
		"username", "johndoe",
		"userid", 123456,
		zap.String("provider", "google"),
	)
}
