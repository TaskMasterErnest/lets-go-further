package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// the showMovieHandler will be used repeatedly
// abstract logic from movies.go into a reusable helper
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// obtain URL parameters stored in the request context - a feature of httprouter
	params := httprouter.ParamsFromContext(r.Context())

	// params are IDs of the string type.
	// convert them into integer types to be used
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid ID parameter")
	}

	return id, nil
}
