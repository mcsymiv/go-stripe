package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mcsymiv/go-stripe/internal/driver"
	"github.com/mcsymiv/go-stripe/internal/web/config"
	"github.com/mcsymiv/go-stripe/internal/web/handlers"
	"github.com/mcsymiv/go-stripe/internal/web/render"
)

const version string = "1.0.0"

var app config.Application

func main() {
	var c config.Config

	// Passed flags to the application
	flag.IntVar(&c.Port, "port", 8082, "Server port")
	flag.StringVar(&c.Env, "env", "dev", "Application environment [ dev | prod ]")
	flag.StringVar(&c.Api, "api", "http://192.168.0.109:8083", "URL to api")
	flag.StringVar(&c.Db.Dsn, "dsn", "mcs:password@tcp(localhost:3306)/db?parseTime=true&tls=false", "DSN")

	flag.Parse()

	// Stripe secret and key
	// pb_ publishable key
	// sk_ secter key is used for Stripe API authentication
	c.Stripe.Secret = os.Getenv("STRIPE_SECRET")
	c.Stripe.Key = os.Getenv("STRIPE_KEY")

	app := &config.Application{
		Config:        &c,
		InfoLog:       log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime),
		ErrorLog:      log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile),
		TemplageCache: make(map[string]*template.Template),
		Version:       version,
	}

	// Returns DB connection pool
	conn, err := driver.OpenDB(c.Db.Dsn)
	if err != nil {
		app.ErrorLog.Printf("unable to get connection pool. Error: %v", err)
		app.ErrorLog.Fatal("connection to DB failed", err)
	}

	defer conn.Close()

	r := handlers.NewRepository(app)
	handlers.New(r)
	render.New(app)

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", c.Port),
		IdleTimeout:       20 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		Handler:           routes(app),
	}

	app.InfoLog.Printf("starting HTTP server in %s mode, on port: %d", c.Env, c.Port)

	err = s.ListenAndServe()
	if err != nil {
		app.ErrorLog.Println("unable to start server", err)
		os.Exit(1)
	}
}
