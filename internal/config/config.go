package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/luxixing/fx-gin/pkg/registry"
	"go.uber.org/fx"
)

func init() {
	registry.Register(
		fx.Options(
			fx.Provide(NewConfig),
		),
	)
}

var (
	once sync.Once
	cfg  Config
)

func NewConfig() (*Config, error) {
	var err error
	once.Do(func() {
		opts := env.Options{
			//
		}
		if err = env.ParseWithOptions(&cfg, opts); err != nil {
			log.Fatalf("failed to parse env: %v", err)
			return
		}
	})
	return &cfg, err
}

type Config struct {
	App      *AppConfig      `env:",init" envPrefix:"APP_"`
	Logger   *LoggerConfig   `env:",init" envPrefix:"LOGGER_"`
	Database *DatabaseConfig `env:",init" envPrefix:"DATABASE_"`
	//todo more
}

type AppConfig struct {
	Name    string `env:"NAME" envDefault:"fx-gin"`
	Version string `env:"VERSION" envDefault:"0.0.1"`
	Host    string `env:"HOST" envDefault:"localhost"`
	Port    int    `env:"PORT" envDefault:"38080"`
	Env     string `env:"ENV" envDefault:"dev"`
}

type DatabaseConfig struct {
	DSN          string `env:"DSN"`
	Database     string `env:"DATABASE" envDefault:"fx-gin.db"`
	ReadTimeout  int    `env:"READ_TIMEOUT" envDefault:"3"`
	WriteTimeout int    `env:"WRITE_TIMEOUT" envDefault:"3"`
	//todo more
}

type LoggerConfig struct {
	Level string `env:"LEVEL" envDefault:"info"`
}

//todo more
