package main

import (
	"net/http"
)

func (app *Config) FetchData(w http.ResponseWriter, r *http.Request) {
	rows, err := SelectQuery(app.Session)

	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
	} else {
		app.writeJson(w, http.StatusOK, rows)
	}

}
