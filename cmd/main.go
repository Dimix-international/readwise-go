package main

import (
	"github.com/Dimix-international/readwise-go/internal/config"
	"github.com/Dimix-international/readwise-go/internal/server"
)

func main() {
	cfg := config.MustLoadConfig()
	server.NewServer(cfg).Run()
}
