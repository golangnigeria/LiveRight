package main

import (
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go LiveRigh api up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllDoctors(w http.ResponseWriter, r *http.Request) {
	doctors, err := app.DB.AllDoctors()
	if err != nil {
		_ = app.errorJSON(w, err)
	}

	_ = app.writeJSON(w, http.StatusOK, doctors)
}

func (app *application) Authenticate(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate against the database

	// Check password
	//
	// create a jwt User
	u := jwtUser{
		ID:        1,
		FirstName: "Prince",
		LastName:  "Dimkpa",
	}

	// generate token
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		_ = app.errorJSON(w, err)
		return
	}

	log.Println(tokens.Token)
	refreshToken := app.auth.GetRefreshToken(tokens.RefreshToken)
	http.SetCookie(w, refreshToken)

	_, _ = w.Write([]byte(tokens.Token))
}
