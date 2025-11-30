package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Some middlerware
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)


	// Farm Routes
	router.HandlerFunc(http.MethodGet, "/api/v1/farms", app.GetFarms)
	router.HandlerFunc(http.MethodPost, "/api/v1/farms", app.InsertFarms)

	// Authentication routes
	router.HandlerFunc(http.MethodPost, "/api/v1/users", app.CreateUser)
	router.HandlerFunc(http.MethodPost, "/api/v1/authenticate", app.Authenticate)
	router.HandlerFunc(http.MethodGet, "/api/v1/refresh", app.refreshToken)

	return app.enableCORS(router)
}


