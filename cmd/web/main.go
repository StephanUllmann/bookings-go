package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/StephanUllmann/bookings-go/pkg/config"
	"github.com/StephanUllmann/bookings-go/pkg/handlers"
	"github.com/StephanUllmann/bookings-go/pkg/render"

	"github.com/alexedwards/scs/v2"
)



const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change to true in prod
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	} 

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)


	render.NewTemplates(&app)
	

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("listening on http://localhost%s\n", portNumber)

	// _ = http.ListenAndServe(portNumber, nil)

	serve := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}