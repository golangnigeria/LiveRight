package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/golangnigeria/MyNneFarm/internal/mailer"
	"github.com/golangnigeria/MyNneFarm/internal/repository"
	"github.com/golangnigeria/MyNneFarm/internal/repository/repo"
	"github.com/joho/godotenv"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string // Data Source Name for the database connection
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	limiter struct {
		enabled bool
		rps     float64
		burst   int
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
	jwt struct {
		JWTIssuer    string
		JWTAudience  string
		JWTSecret    string
		CookieDomain string
	}
}

type application struct {
	config config
	logger *log.Logger
	DB     repository.DatabaseRepository
	mailer mailer.Mailer
	wg     sync.WaitGroup
	auth   Auth
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "dsn", os.Getenv("DATABASE_URL"), "PostgreSQL data source name")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP server host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 2525, "SMTP server port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "4ede3760dd00c6", "SMTP server username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "1f87de4dec9bdb", "SMTP server password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "stackninja.pro@gmail.com", "SMTP sender email address")

	// Rate limiting configuration
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiting")
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Requests per second for rate limiting")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Burst size for rate limiting")

	// JWT configuration
	flag.StringVar(&cfg.jwt.JWTSecret, "jwt-secret", "verysecret", "Signing secret for our help text")
	flag.StringVar(&cfg.jwt.JWTIssuer, "jwt-Issuer", "example.com", "Signing Issuer for our help text")
	flag.StringVar(&cfg.jwt.JWTAudience, "jwt-Audience", "example.com", "Signing Audience for our help text")
	flag.StringVar(&cfg.jwt.CookieDomain, "Cookie-Domain", "localhost", "Domain for development environtment")

	// Parse command line flags
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config: cfg,
		logger: logger,
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	conn, err := app.connectToDB()
	if err != nil {
		app.logger.Fatal(err)
	}

	app.DB = &repo.NeonDBRepo{DB: conn}
	defer app.DB.Connection().Close()


	// JWT populating the default setting for the jwt.
	app.auth = Auth{
		Issuer: cfg.jwt.JWTIssuer,
		Audience: cfg.jwt.JWTAudience,
		Secret: cfg.jwt.JWTSecret,
		TokenExpiry: time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath: "/",
		CookieName: "__Host-refresh_token",
		CookieDomain: cfg.jwt.CookieDomain,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Printf("Starting %s server on port %d", app.config.env, app.config.port)
	err = srv.ListenAndServe()
	if err != nil {
		app.logger.Fatal(err)
	}
}
