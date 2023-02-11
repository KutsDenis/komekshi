package cfg

import (
	"github.com/ilyakaznacheev/cleanenv"
	"komekshi/pkg/logger"
	"sync"
)

type Config struct {
	Token string `yaml:"token"`
}

var Get *Config
var onceCfg sync.Once

func Load(l *logger.Logger) {
	onceCfg.Do(func() {
		Get = &Config{}

		if err := cleanenv.ReadConfig("configs/config.yml", Get); err != nil {
			l.Fatal(err)
		}
	})
}
