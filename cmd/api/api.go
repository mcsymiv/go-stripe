package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		key    string
		secret string
	}
}

type app struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templageCache map[string]*template.Template
	version       string
}

func (a *app) serve() error {
	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", a.config.port),
		IdleTimeout:       20 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		Handler:           a.routes(),
	}

	a.infoLog.Printf("starting BE server in %s mode, on port: %d", a.config.env, a.config.port)

	return s.ListenAndServe()
}

func main() {
	var c config

	flag.IntVar(&c.port, "port", 8083, "Server API port")
	flag.StringVar(&c.env, "env", "dev", "Application environment [ dev | prod | maintenance ]")

	flag.Parse()

	c.stripe.secret = os.Getenv("STRIPE_SECRET")
	c.stripe.key = os.Getenv("STRIPE_KEY")

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &app{
		config:   c,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	err := app.serve()
	if err != nil {
		app.errorLog.Println("unable to start server", err)
		os.Exit(1)
	}
}
