package main

import (
	"flag"
	"log"
	"server/internal/server"
	"server/package/config"
)

func main() {
	configPath := flag.String("c", "", "config file path")
	flag.Parse()

	if *configPath == "" {
		log.Fatal("require config file path")
	}
	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("err while reading config file: %v", err)
	}
	server.Serve(&cfg)
}
