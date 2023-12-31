package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/snipersap/bookings/pkg/config"
	"github.com/snipersap/bookings/pkg/handlers"
	"github.com/snipersap/bookings/pkg/render"
)

const portNumber string = ":8080"

// declare app config
var app config.AppConfig

func main() {

	s := CreateSession()
	SetSession(&app, s)

	//create the Template Cache
	tc, err := render.GetTemplateCache()
	if err != nil {
		log.Fatalln("cannot load template cache. Exiting app:", err.Error())
	}

	// Set the template cache to app config and render package
	SetTmplCache(&app, tc)
	SetUseCache(&app, false)
	render.SetAppConfig(&app)
	// Init a handler repository with app config
	repo := handlers.NewRepo(&app)
	handlers.SetRepo(repo)

	// Setup server and run with pat
	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	log.Println("Starting web server on port:", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln("couldn't start the web server on port "+portNumber+
			" Error was:", err.Error())
	}
}

// CreateSession creates a new session with default values and returns it
func CreateSession() *scs.SessionManager {
	s := scs.New()
	s.Lifetime = 24 * time.Hour
	s.Cookie.Persist = true
	s.Cookie.SameSite = http.SameSiteLaxMode
	s.Cookie.Secure = app.InProduction
	return s
}

// SetSession sets the provided session to the app config
func SetSession(app *config.AppConfig, s *scs.SessionManager) {
	app.Session = s
}

// SetUseCache sets the provided value for UseCache to app config
func SetUseCache(app *config.AppConfig, uc bool) {
	app.UseCache = uc
}

// SetTmplCache sets the template cache provided to app config
func SetTmplCache(app *config.AppConfig, t map[string]*template.Template) {
	app.TemplateCache = t
}
