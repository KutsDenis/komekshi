package main

import (
	"komekshi/internal/cfg"
	"komekshi/internal/discord"
	"komekshi/pkg/logger"
)

func main() {
	l := logger.GetLogger()
	cfg.Load(&l)

	discord.Start(&l)

	<-make(chan struct{})
	return
}
