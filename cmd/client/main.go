package main

import (
	"context"
	"flag"
	"log"

	"github.com/theoptz/grpcchat/internal/sigctx"
	"github.com/theoptz/grpcchat/pkg/client"
)

var (
	addr  string
	name  string
	debug bool
)

func init() {
	flag.StringVar(&addr, "a", "0.0.0.0:8000", "server address")
	flag.StringVar(&name, "n", "joe", "user name")
	flag.BoolVar(&debug, "d", false, "debug mode")

	flag.Parse()
}

func main() {
	cl, err := client.New(addr, name, debug)
	if err != nil {
		log.Fatal(err)
	}

	ctx := sigctx.New(context.Background())

	err = cl.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
