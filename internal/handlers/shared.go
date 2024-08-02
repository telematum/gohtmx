package handlers

import (
	"log/slog"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) error

func Make(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("error serving", "path", slog.StringValue(r.URL.Path), "reason", err)
		}
	}
}
