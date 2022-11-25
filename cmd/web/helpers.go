package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Writes an error message and stack trace to the errorLog, then send a generic
// 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

// Sends a specific status code and corresponding description to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Sends a 404 Not Found Respoonse to the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// Renders a page if not render before add to templateCache.
func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}

}
