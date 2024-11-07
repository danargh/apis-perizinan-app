package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/danargh/apis-perizinan-app/internal"
	"github.com/danargh/apis-perizinan-app/pkg/database"
	"github.com/danargh/apis-perizinan-app/pkg/version"

	"github.com/lmittmann/tint"
)

func main() {
	// setup log and coloring log with tint
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

// running the app
func run(logger *slog.Logger) error {
	cfg, err := internal.NewConfig()

	// flag.Bool() menghasilkan nilai pointer sehingga untuk dapat nilainya perlu *showVersion
	showVersion := flag.Bool("version", false, "display version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	// close db connection saat fungsi run hampir selesai menjalankan semua baris code
	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)
	if err != nil {
		return err
	}
	defer db.Close()

	// membuat pointer app dari type application
	app := &Application{
		config: cfg,
		db:     db,
		logger: logger,
	}

	return app.serveHTTP()
}
