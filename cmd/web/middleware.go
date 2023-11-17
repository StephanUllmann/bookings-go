package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		currTime := time.Now()
		dt := fmt.Sprint(currTime.Format("02.01.2006 15:04:05"))

		fmt.Printf("Hit the page: %v | %v | %v\n", r.Method, r.URL.Path, dt)
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// The SessionLoad function is a middleware that loads and saves session data for an HTTP handler.
func SessionLoad (next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}