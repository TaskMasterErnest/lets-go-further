package main

import (
	"fmt"
	"net/http"
)

// handler for creating a movie
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "creating a new movie")
}

// handler for showing movies based on ID requested
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// use the helper method for obtaining ID
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// placeholder response
	fmt.Fprintf(w, "show details of movie %d\n", id)
}
