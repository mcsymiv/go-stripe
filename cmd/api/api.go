package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mcsymiv/go-stripe/internal/api/config"
	"github.com/mcsymiv/go-stripe/internal/api/handlers"
)

const version = "1.0.0"

var app config.Application

func main() {
	var c config.Config

	flag.IntVar(&c.Port, "port", 8083, "Server API port")
	flag.StringVar(&c.Env, "env", "dev", "Application environment [ dev | prod | maintenance ]")

	flag.Parse()

	c.Stripe.Secret = os.Getenv("STRIPE_SECRET")
	c.Stripe.Key = os.Getenv("STRIPE_KEY")

	app := &config.Application{
		Config:   &c,
		InfoLog:  log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile),
		Version:  version,
	}

	repository := handlers.NewRepository(app)
	handlers.NewHandlers(repository)

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", c.Port),
		IdleTimeout:       20 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		Handler:           routes(app),
	}

	app.InfoLog.Printf("starting BE server in %s mode, on port: %d", c.Env, c.Port)

	err := s.ListenAndServe()
	if err != nil {
		app.ErrorLog.Println("unable to start server", err)
		os.Exit(1)
	}
}
