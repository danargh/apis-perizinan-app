package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/danargh/apis-perizinan-app/internal/database"
	"github.com/danargh/apis-perizinan-app/internal/env"
	"github.com/danargh/apis-perizinan-app/internal/version"

	"github.com/lmittmann/tint"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	baseURL   string
	httpPort  int
	basicAuth struct {
		username       string
		hashedPassword string
	}
	cookie struct {
		secretKey string
	}
	db struct {
		dsn         string
		automigrate bool
	}
	jwt struct {
		secretKey string
	}
}

type application struct {
	config config
	db     *database.DB
	logger *slog.Logger
	wg     sync.WaitGroup
}

func run(logger *slog.Logger) error {
	var cfg config

	cfg.baseURL = env.GetString("BASE_URL", "http://localhost:4444")
	cfg.httpPort = env.GetInt("HTTP_PORT", 4444)
	cfg.basicAuth.username = env.GetString("BASIC_AUTH_USERNAME", "admin")
	cfg.basicAuth.hashedPassword = env.GetString("BASIC_AUTH_HASHED_PASSWORD", "$2a$10$jRb2qniNcoCyQM23T59RfeEQUbgdAXfR6S0scynmKfJa5Gj3arGJa")
	cfg.cookie.secretKey = env.GetString("COOKIE_SECRET_KEY", "osxcxvka3ke4335avd6zpaf2lnkdwk5t")
	cfg.db.dsn = env.GetString("DB_DSN", "postgres:180701@localhost:5432/db_apis_perizinan?sslmode=disable")
	cfg.db.automigrate = env.GetBool("DB_AUTOMIGRATE", true)
	cfg.jwt.secretKey = env.GetString("JWT_SECRET_KEY", "sfxfvu37244ljpzpmfmqvvgcft6d33kb")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)
	if err != nil {
		return err
	}
	defer db.Close()

	app := &application{
		config: cfg,
		db:     db,
		logger: logger,
	}

	return app.serveHTTP()
}