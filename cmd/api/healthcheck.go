package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":  "available",
		"env":     app.config.env,
		"version": version,
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"healthcheck": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
