package config

import "github.com/caarlos0/env/v10"

type Env struct {
	// DB
	DBType string `env:"DB_TYPE,required" envDefault:"mysql"`
	DBHost string `env:"DB_HOST,required" envDefault:"db"`
	DBPort string `env:"DB_PORT,required" envDefault:"13306"`
	DBUser string `env:"DB_USER,required" envDefault:"root"`
	DBPass string `env:"DB_PASS,required" envDefault:""`
	DBName string `env:"DB_NAME,required" envDefault:"pwk"`
	DBLoc  string `env:"DB_LOC,required" envDefault:"Asia%2FTokyo"`
}

func LoadEnv() (*Env, error) {
	var conf Env
	if err := env.Parse(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
