package config

import (
  "fmt"
  "github.com/caarlos0/env/v6"
)

type Config struct {
  Port string `env:"PORT" envDefault:":8080"`
  Dsn  string `env:"DSN"`
}

var cfg Config

func Init() {
  if err := env.Parse(&cfg); err != nil {
    fmt.Printf("%+v\n", err)
  }

  fmt.Printf("%+v\n", cfg)
}
func GetConfig() *Config {
  return &cfg
}
