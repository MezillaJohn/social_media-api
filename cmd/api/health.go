package main

import (
	"net/http"
)

func (app *application) handleApiHealthCheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	if err := writeJson(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
