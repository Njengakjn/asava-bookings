package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

type config struct {
	port         int
	env          string
	usecache     bool
	inProduction bool
	db           struct {
		dsn         string
		maxIdleTime time.Duration
	}
}

type templateData struct {
	SMap      map[string]string
	IMap      map[string]int
	FMap      map[string]float32
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

type application struct {
	tc      map[string]*template.Template
	config  config
	session *scs.SessionManager
}

func main() {
	var cfg config
	cfg.inProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = cfg.inProduction

	flag.IntVar(&cfg.port, "port", 8080, "application port")
	flag.Parse()

	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		tc:      tc,
		config:  cfg,
		session: session,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      routes(app),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	srv.ListenAndServe()
}
