package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	models "github.com/golangnigeria/MyNneFarm/internal/model"
)

// GetFarms handles the HTTP request to retrieve all farms.
// It uses the NeonDBRepo to fetch the farms and writes the response in JSON format.
func (app *application) GetFarms(w http.ResponseWriter, r *http.Request) {
	// This is a placeholder for the actual implementation of the GetFarms handler.
	farms, err := app.DB.GetFarms()
	if err != nil {
		app.logger.Println("Error retrieving farms:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"farms": farms}, nil)
	if err != nil {
		app.logger.Println("Error writing JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// InsertFarms handles the HTTP request to insert a new farm.
func (app *application) InsertFarms(w http.ResponseWriter, r *http.Request) {

	var farm models.FarmModel

	// Read the incoming JSON into the farm model
	err := app.readJSON(w, r, &farm)
	if err != nil {
		app.logger.Println("Error reading JSON request:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Validate the incoming data
	if err := farm.Validate(); err != nil {
		app.logger.Println("Error validating farm:", err)
		http.Error(w, "Validation Error", http.StatusBadRequest)
		return
	}

	// Save to DB
	farmID, err := app.DB.InsertFarm(farm)
	if err != nil {
		app.logger.Println("Error inserting farm:", err)
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	farm.ID = int(farmID)
	w.Header().Set("Location", fmt.Sprintf("/farms/%d", farm.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"farm": farm}, nil)
	if err != nil {
		app.logger.Println("Error writing JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// CreateUser handles the HTTP request to create a new user.
func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FullName string `json:"full_name" db:"full_name"`
		Email    string `json:"email" db:"email"`
		Phone    string `json:"phone" db:"phone"`
		Password string `json:"-" db:"password"`
	}

	// Read the incoming JSON into the user model
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "unable to read the create user request payload")
		return
	}

	user := &models.User{
		FullName:  input.FullName,
		Email:     input.Email,
		Phone:     input.Phone,
		Activated: false,              // Default to false
		Roles:     []string{"farmer"}, // Default role
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		log.Println("Password Set error:", err)
		app.errorResponse(w, r, http.StatusUnauthorized, "invalid email or password")
		return
	}

	// Save to DB
	err = app.DB.CreateUser(user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// send a welcome email

	// Writing out the json
	err = app.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil)
	if err != nil {
		app.logger.Println("Error writing JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Authenticate handles the HTTP request to validate the user credentials
func (app *application) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid request payload")
		return
	}

	log.Printf("Incoming login: %s / %s", requestPayload.Email, requestPayload.Password)

	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		log.Println("DB error:", err)
		app.errorResponse(w, r, http.StatusUnauthorized, "invalid email or password")
		return
	}

	err = user.Password.Set(requestPayload.Password)
	if err != nil {
		log.Println("Password Set error:", err)
		app.errorResponse(w, r, http.StatusUnauthorized, "invalid email or password")
		return
	}

	log.Println("Fetched hashed password:", user.Password)

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil {
		log.Println("Password comparison error:", err)
		app.errorResponse(w, r, http.StatusInternalServerError, "error checking password")
		return
	}
	if !valid {
		log.Println("Invalid password")
		app.errorResponse(w, r, http.StatusUnauthorized, "invalid email or password")
		return
	}

	u := jwtUser{
		ID:       user.ID,
		FullName: user.FullName,
	}

	tokens, err := app.auth.GenerateTokenPairs(&u)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, "token generation failed")
		return
	}

	http.SetCookie(w, app.auth.GetRefreshCookie(tokens.RefreshToken))
	app.writeJSON(w, http.StatusAccepted, envelope{"user": tokens}, nil)
}

// Refresh handles the HTTP request to refresh a user
func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			// Parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(app.auth.Secret), nil
			})
			if err != nil {
				app.errorResponse(w, r, http.StatusUnauthorized, "Unautherized")
				return
			}

			// get the user ID from the claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorResponse(w, r, http.StatusUnauthorized, "Unknown user")
				return
			}

			user, err := app.DB.GetUserByID(userID)
			if err != nil {
				app.errorResponse(w, r, http.StatusUnauthorized, "UNKNOWN USER")
				return
			}

			u := jwtUser{
				ID:       user.ID,
				FullName: user.FullName,
			}

			// Generate a token pair for this user
			tokenPairs, err := app.auth.GenerateTokenPairs(&u)
			if err != nil {
				app.errorResponse(w, r, http.StatusUnauthorized, "Error generating tokens")
				return
			}

			// Set the refresh token as cookie on the user browser
			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))

			// WriteJSON  write out the out put to the developer
			app.writeJSON(w, http.StatusAccepted, envelope{
				"Tokens":tokenPairs,
				"RefreshToken": tokenPairs.RefreshToken,
			 }, nil)
		}
	}
}
