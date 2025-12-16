package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/golangnigeria/liveright_backend/internal/models"
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
	userEmail, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		_ = app.errorJSON(w, errors.New("invalid crediantial"), http.StatusBadRequest)
		return
	}

	// Check password
	valid, err := userEmail.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		_ = app.errorJSON(w, errors.New("incorrect password"), http.StatusBadRequest)
		return
	}

	// create a jwt User
	u := jwtUser{
		ID:        userEmail.ID,
		FirstName: userEmail.FirstName,
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

	_ = app.writeJSON(w, http.StatusAccepted, map[string]any{
		"message": "welcome back " + userEmail.FirstName + ", This is LiveRight.",
		"user": map[string]any{
			"name":   userEmail.FirstName + " " + userEmail.LastName,
			"email":  userEmail.Email,
			"active": userEmail.Active,
		},
		"tokens": tokens,
	})
}

// RegisterPatient
func (app *application) RegisterPatient(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Phone     string `json:"phone,omitempty"`
	}

	if err := app.readJSON(w, r, &payload); err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// check existing
	_, err := app.DB.GetUserByEmail(payload.Email)
	if err == nil {
		_ = app.errorJSON(w, errors.New("email already registered"), http.StatusBadRequest)
		return
	}

	u := &models.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     models.Email(payload.Email),
		Phone:     &payload.Phone,
		Active:    true,
		RoleID:    models.Role{ID: 1}, // patient
	}
	if err := u.HashPassword(payload.Password); err != nil {
		_ = app.errorJSON(w, errors.New("unable to hash password"), http.StatusInternalServerError)
		return
	}

	newUser, err := app.DB.InsertUser(u)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	j := jwtUser{
		ID:        newUser.ID,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
	}

	tokens, err := app.auth.GenerateTokenPair(&j)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, app.auth.GetRefreshToken(tokens.RefreshToken))

	_ = app.writeJSON(w, http.StatusCreated, map[string]any{
		"message": "patient registered",
		"user": map[string]any{
			"id":         newUser.ID,
			"first_name": newUser.FirstName,
			"email":      newUser.Email,
		},
		"tokens": tokens,
	})
}
