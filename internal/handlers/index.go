package handlers

import (
	"fmt"
	"net/http"

	"cloudevents-explorer/internal/templates"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Only handle exact "/" path to avoid catching all routes
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, templates.Index)
}
