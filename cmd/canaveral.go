package main

import (
	"fmt"
	"log"

	"github.com/canaveral/app"
	"github.com/canaveral/cli"
	"github.com/canaveral/config"
	// for demo
)

func main() {
	cfg, err := config.Init("config", "config", "toml")
	if err != nil {
		log.Fatal(fmt.Errorf("config initialization error: %w", err))
	}
	a, err := app.New(cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("app initialization error: %w", err))
	}
	cli := cli.NewRootCmd(a)
	cli.Run()
	defer a.Registry.Close()
}
