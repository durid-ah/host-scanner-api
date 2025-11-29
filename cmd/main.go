package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/durid-ah/host-scanner-api/config"
	cronscheduler "github.com/durid-ah/host-scanner-api/cron_scheduler"
	"github.com/durid-ah/host-scanner-api/db"
	"github.com/durid-ah/host-scanner-api/handler"
	nmapscanner "github.com/durid-ah/host-scanner-api/scanner"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
)

func createFiberServer(storage *db.Storage) *fiber.App {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Nmap API", "0.0.1"))

	huma.Get(api, "/api/v1/hosts", handler.GetAllHosts(storage))
	huma.Get(api, "/api/v1/hosts/{hostname}", handler.GetHost(storage))
	return app
}

func main() {
	cfg := config.NewConfig()

	storage, err := db.NewStorage(slog.Default())
	if err != nil {
		log.Fatal(err)
	}

	opts := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewJSONHandler(os.Stdout, &opts)
	slog.SetDefault(slog.New(handler))

	// run the scanner once at startup to populate the db
	slog.Info("running intial scan to populate the db...")
	scanTask := nmapscanner.CreateScannerTask(*storage, cfg)
	scanTask(context.Background())
	slog.Info("intial scan completed")

	scheduler := cronscheduler.NewBackgroundScheduler(*storage, cfg)
	scheduler.Start()

	defer func() {
		slog.Info("shutting down scheduler")
		err := scheduler.Shutdown()
		if err != nil {
			slog.Error("error shutting down scheduler", "error", err)
			log.Fatal(err)
		}
	}()

	app := createFiberServer(storage)
	app.Listen(fmt.Sprintf("%s:%s", cfg.NmapAPIHost, cfg.NmapAPIPort))
}
