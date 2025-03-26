package main

import "net/http"

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		mehtod = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", mehtod, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
