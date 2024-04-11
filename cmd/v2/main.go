package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/bytedance/gopkg/util/gctuner"
	"github.com/corentings/chessbet/app"
	"github.com/corentings/chessbet/cmd/v2/config"
	"github.com/corentings/chessbet/infrastructures"
	"github.com/corentings/chessbet/pkg/crypto"
	"github.com/corentings/chessbet/pkg/jwt"
	"github.com/corentings/chessbet/pkg/logger"
	"github.com/corentings/chessbet/pkg/oauth"
	"github.com/pkg/errors"
)

func main() {
	configPath := config.GetConfigPath(config.IsDevelopment())

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("❌ Error loading config: %s", err.Error())
	}

	logger.GetLogger().SetLogLevel(cfg.Log.GetSlogLevel()).CreateGlobalHandler()

	setup(cfg)

	e := app.CreateEchoInstance(cfg.Server)

	if cfg.Log.GetSlogLevel().Level() == slog.LevelDebug {
		slog.Debug("🔧 Debug mode enabled")
	}

	slog.Info("starting server 🚀", slog.String("version", cfg.Server.AppVersion))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err = e.Start(); err != nil {
			slog.Error("error starting server", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	const shutdownTimeout = 10 * time.Second

	<-ctx.Done()
	_, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	slog.Info("shutting down server")

	if err = shutdown(); err != nil {
		slog.Error("error shutting down server", slog.Any("error", err))
	}

	slog.Info("server stopped")
}

func setup(cfg *config.Config) {
	setupInfrastructures(cfg)

	gcTuning()

	setupOAuth(cfg)

	setupJwt(cfg)

	slog.Info("✅ setup completed!")
}

func setupJwt(cfg *config.Config) {
	// Parse the keys
	if err := crypto.GetKeyManagerInstance().ParseEd25519Key(); err != nil {
		log.Fatal("❌ Error parsing keys", slog.Any("error", err))
	}

	// Create the JWT instance
	jwtInstance := jwt.NewJWTInstance(cfg.Auth.JWTHeaderLen, cfg.Auth.JWTExpiration,
		crypto.GetKeyManagerInstance().GetPublicKey(), crypto.GetKeyManagerInstance().GetPrivateKey())

	jwt.GetJwtInstance().SetJwt(jwtInstance)

	slog.Info("✅ Created JWT instance")
}

func shutdown() error {
	e := app.GetEchoInstance()

	slog.Info("🔒 Server shutting down...")

	var shutdownError error

	if err := e.Shutdown(context.Background()); err != nil {
		// Add the error to the shutdownError
		shutdownError = errors.Wrap(shutdownError, err.Error())
	}

	slog.Info("🧹 Running cleanup tasks...")

	infrastructures.GetPgxConnInstance().ClosePgx()

	slog.Info("✅ Disconnected from database")

	slog.Info("✅ Cleanup tasks completed!")

	return shutdownError
}

func setupInfrastructures(cfg *config.Config) {
	err := infrastructures.NewPgxConnInstance(infrastructures.PgxConfig{
		DSN: cfg.Pgx.DSN,
	}).ConnectPgx()
	if err != nil {
		log.Fatal("❌ Error connecting to database", slog.Any("error", err))
	}

	slog.Info("✅ Connected to database")
}

func setupOAuth(cfg *config.Config) {
	oauthConfig := oauth.GlobalConfig{
		CallbackURL: cfg.Server.Host,
		FrontendURL: cfg.Server.FrontendURL,
	}

	oauth.SetOauthConfig(oauthConfig)

	oauth.InitGithub(cfg.Auth.Github)
	oauth.InitDiscord(cfg.Auth.Discord)

	slog.Info("✅ OAuth setup completed!")
}

func gcTuning() {
	var limit float64 = 4 * config.GCLimit
	// Set the GC threshold to 70% of the limit
	threshold := uint64(limit * config.GCThresholdPercent)

	gctuner.Tuning(threshold)

	slog.Info(fmt.Sprintf("🔧 GC Tuning - Limit: %.2f GB, Threshold: %d bytes, GC Percent: %d, Min GC Percent: %d, Max GC Percent: %d",
		limit/(config.GCLimit),
		threshold,
		gctuner.GetGCPercent(),
		gctuner.GetMinGCPercent(),
		gctuner.GetMaxGCPercent()))

	slog.Info("✅ GC Tuning completed!")
}
