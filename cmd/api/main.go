package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangnigeria/liveright_backend/internal/repository"
	"github.com/golangnigeria/liveright_backend/internal/repository/dbrepo"
	"github.com/joho/godotenv"
)

const port = 8080

type application struct {
	Domain       string
	DSN          string
	DB           repository.DatabaseRepo
	auth         Auth
	JwtSecret    string
	JWTIssuer    string
	JWTAudienc   string
	CookieDomain string
}

func main() {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using defaults and flags")
	}

	var app application

	// Set flags, default to environment variables if they exist
	flag.StringVar(&app.DSN, "dsn", os.Getenv("DSN"), "Postgres DSN")
	flag.StringVar(&app.JwtSecret, "jwt-secret", os.Getenv("JWT_SECRET"), "Signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", os.Getenv("JWT_ISSUER"), "Signing issuer")
	flag.StringVar(&app.JWTAudienc, "jwt-audience", os.Getenv("JWT_AUDIENCE"), "Signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", os.Getenv("COOKIE_DOMAIN"), "Cookie domain")
	flag.StringVar(&app.Domain, "domain", os.Getenv("DOMAIN"), "Domain")

	flag.Parse()

	// Connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	defer func() {
		err := app.DB.Connection().Close()
		if err != nil {
			fmt.Println("Error closing DB", err)
		}
	}()

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudienc,
		Secret:        app.JwtSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  app.Domain,
	}

	log.Println("Starting application on port", port)

	// Start the web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
