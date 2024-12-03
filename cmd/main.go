package main

import (
	"log"
	"passport/internal/srvenv"
)

func main() {
	cfg, err := srvenv.ReadConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	server, err := srvenv.NewServer(cfg)
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	server.Run(cfg.Server.Host + ":" + cfg.Server.Port)
}
