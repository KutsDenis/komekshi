package main

import (
	"komekshi/internal/broker"
	"komekshi/internal/cfg"
	"komekshi/internal/discord"
	"komekshi/pkg/logger"
)

func main() {
	l := logger.GetLogger()
	cfg.Load(&l)

	discord.Start(&l)
	go broker.Broker()
	<-make(chan struct{})
	return
}
