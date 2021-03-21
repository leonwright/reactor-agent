package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/leonwright/devhelper/pkg/config"
	"github.com/leonwright/devhelper/pkg/network"
)

func run(ctx context.Context, c *config.Config, out io.Writer) error {
	c.Init(os.Args)
	log.SetOutput(out)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(c.DNSUpdateInterval):
			var ip string = network.GetCurrentIP()
			network.UpdateDevDNS(c, ip)
		}
	}
}

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting DevHelper Daemon...\n")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case msg := <-signalChan:
			log.Printf("Got %s message, exiting.", msg.String())
			cancel()
			os.Exit(1)
		case <-ctx.Done():
			log.Printf("Done.")
			os.Exit(1)
		}
	}()

	c := &config.Config{}

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	if err := run(ctx, c, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
