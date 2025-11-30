package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func openDB(dsn string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	// Open a connection to the database
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}

	// Set database connection parameters
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	maxIdleDuration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, fmt.Errorf("cannot parse max idle time: %w", err)
	}

	db.SetConnMaxIdleTime(maxIdleDuration)

	// Create a context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping the database to ensure the connection is established
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.config.db.dsn, app.config.db.maxOpenConns, app.config.db.maxIdleConns, app.config.db.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	app.logger.Printf("Database connection established successfully")
	return connection, nil
}
