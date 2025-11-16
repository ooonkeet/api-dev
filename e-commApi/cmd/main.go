package main

import (
	// "log"
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}
	api := application{
		config: cfg,
	}
	logger:=slog.New(slog.NewTextHandler(os.Stdout,nil))
	slog.SetDefault(logger)
	// structured logging
	if err := api.run(api.mount()); err != nil {
		slog.Error("server has failed to start, err: %s", err)
		os.Exit(1)
	}
}
