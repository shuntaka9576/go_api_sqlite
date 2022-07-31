package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/shuntaka9576/go_api_sqlite/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("faild to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("faild to listen port %d: %v", cfg.Port, err)

	}
	log.Printf("sqlite db path: %s", cfg.DBPath)

	mux, cleanup, err := NewMux(ctx, cfg)
	if err != nil {
		return err
	}
	s := NewServer(l, mux)
	defer cleanup()

	return s.Run(ctx)
}
