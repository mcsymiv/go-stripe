package config

import (
	"html/template"
	"log"
)

type Config struct {
	Port int
	Env  string
	Api  string
	Db   struct {
		dsn string
	}
	Stripe struct {
		Key    string
		Secret string
	}
}

type Application struct {
	Config        *Config
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	TemplageCache map[string]*template.Template
	Version       string
}
