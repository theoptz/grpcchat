package main

import (
	"context"
	"flag"
	"log"

	"github.com/theoptz/grpcchat/internal/sigctx"
	"github.com/theoptz/grpcchat/pkg/server"
)

var (
	addr  string
	debug bool
)

func init() {
	flag.StringVar(&addr, "a", "0.0.0.0:8000", "server address")
	flag.BoolVar(&debug, "d", false, "debug mode")

	flag.Parse()
}

func main() {
	srv, err := server.New(addr, debug)
	if err != nil {
		log.Fatal(err)
	}

	ctx := sigctx.New(context.Background())

	err = srv.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
