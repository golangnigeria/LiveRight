package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// Marshal the data into JSON format with indentation for readability
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
 
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		 
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		 
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		 
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
 
		default:
			return err
		}
	}
	return nil
}


func (app *application) background(fn func()) {
	app.wg.Add(1)

	go func() {

		defer app.wg.Done()
		 // Recover from any panic that occurs in the background function
		defer func() {
			if r := recover(); r != nil {
				app.logger.Println("Background task panicked:", r)
			}
		}()

		fn()
	}()
}