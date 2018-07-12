package main

import (
	"net/http"

	"github.com/prest/cmd"
	"github.com/prest/config"
	"github.com/prest/config/router"
)

func main() {
	config.Load()

	// Get pPREST router
	r := router.Get()

	// Register custom routes
	r.HandleFunc("/ping", Pong).Methods("GET")
	r.HandleFunc("/{database:demo}/{schema:public}/{table:person}", overrideHandler).Methods("GET")

	// Call pREST cmd
	cmd.Execute()
}

// Pong is a healthcheck handler
func Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong!"))
}

func overrideHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nothing to see here!"))
}
