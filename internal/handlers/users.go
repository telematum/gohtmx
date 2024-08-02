package handlers

import (
	"net/http"

	"github.com/telematum/gohtmx/internal/views"

	"github.com/telematum/gohtmx/internal/bll"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) error {
	data, err := bll.GetUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return views.ListUsers(data).Render(r.Context(), w)
}
