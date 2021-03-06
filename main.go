package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prest/cmd"
	"github.com/prest/config"
	"github.com/prest/config/router"
	"github.com/urfave/negroni"
)

func main() {
	config.Load()

	// Get pPREST router
	r := router.Get()

	// Register custom routes
	r.HandleFunc("/ping", Pong).Methods("GET")
	// handler overload
	r.HandleFunc("/{database:demo}/{schema:public}/{table:person}", OverloadedHandler).Methods("GET")

	// custom middleware applied in just one endpoint
	adminRoutes := mux.NewRouter().PathPrefix("/admin").Subrouter()
	adminRoutes.HandleFunc("/secret", SecretHandler)
	r.PathPrefix("/admin").Handler(negroni.New(
		negroni.HandlerFunc(adminOnly),
		negroni.Wrap(adminRoutes),
	))

	// owner secret is just like admin but using only gorilla
	ownerRoutes := r.PathPrefix("/owner").Subrouter()
	ownerRoutes.HandleFunc("/secret", SecretHandler)
	ownerRoutes.Use(ownerOnly)

	// Call pREST cmd
	cmd.Execute()
}

// Pong is a healthcheck handler
func Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong!"))
}

// OverloadedHandler just change prest default endpoint
func OverloadedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nothing to see here!"))
}

// SecretHandler hides a secret
func SecretHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("import this!"))
}

func adminOnly(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Header.Get("X-user") != "admin" {
		http.Error(w, "nothing to see here", http.StatusUnauthorized)
		return
	}
	next(w, r)
}

func ownerOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-user") != "owner" {
			http.Error(w, "nothing to see here", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
