package main

import "net/http"

func (app *application) handleApiHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))

}
