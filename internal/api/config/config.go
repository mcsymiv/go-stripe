package config

import "log"

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
	Stripe struct {
		Key    string
		Secret string
	}
}

type Application struct {
	Config   *Config
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Version  string
}
