package handlers

import (
	"net/http"

	"github.com/telematum/gohtmx/internal/views"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return views.Home().Render(r.Context(), w)
}
