package main

import (
	// "log"
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/ooonkeet/api-dev/e-commApi/internal/env"
)

func main() {
	ctx:=context.Background()
	cfg := config{
		addr: ":8080",
		db:   dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING","host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}
	// logger
	logger:=slog.New(slog.NewTextHandler(os.Stdout,nil))
	slog.SetDefault(logger)
	// database
	conn,err := pgx.Connect(ctx, cfg.db.dsn)
	if err!=nil{
		panic(err)
	}
	defer conn.Close(ctx)
	logger.Info("connected to database","dsn",cfg.db.dsn)
	api := application{
		config: cfg,
		db: conn,
	}
	
	// structured logging
	if err := api.run(api.mount()); err != nil {
		slog.Error("server has failed to start, err: %s", err)
		os.Exit(1)
	}
}
