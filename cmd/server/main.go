package main

import (
	"cmp"
	"log/slog"
	"net/http"
	"os"

	"github.com/telematum/gohtmx/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	serv := mux.NewRouter()

	serv.HandleFunc("/", handlers.Make(handlers.HandleHome))
	serv.HandleFunc("/users", handlers.Make(handlers.UsersHandler))

	listenAddr := cmp.Or(os.Getenv("LISTENADDR"), ":4000")
	slog.Info("starting server", "address", listenAddr)
	if err := http.ListenAndServe(listenAddr, serv); err != nil {
		slog.Error("server bind failed", "listen", slog.StringValue(listenAddr), "error", err)
		os.Exit(1)
	}
}
