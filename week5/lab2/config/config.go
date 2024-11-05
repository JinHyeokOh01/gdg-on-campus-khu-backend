package config

import(
	"github.com/caarlos0/env/v6"
)

type Config struct{
	Env string `env:"TODO_ENV" envDefault:"Dev"`
	Port int `env:"PORT" envDefault:"80"`
}

func New()(*Config, error){
	cfg := &config{}
	if err := env.Parse(cfg); err != nil{
		return nil, err
	}
	return cfg, nil
}