package broker

import (
	"komekshi/internal/discord"
	"komekshi/internal/discord/events"
	"komekshi/pkg/logger"
	"komekshi/pkg/scheduler"
	"time"
)

func Broker() {
	updateEvents := scheduler.Timer{Min: 1}

	go updateEvents.EachTime(checkEventTime)
}

func checkEventTime() {
	l := logger.GetLogger()

	e := events.List(discord.Bot)

	for i := 0; i < len(e); i++ {
		if e[i].ScheduledStartTime.Before(time.Now().UTC().Add(75*time.Second)) && e[i].Status == 1 {
			events.Start(discord.Bot, e[i], &l)
			l.Info("Starting " + e[i].Name)
		}
	}
}
