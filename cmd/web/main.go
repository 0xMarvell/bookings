package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Marvellous-Chimaraoke/bookings/internal/config"
	"github.com/Marvellous-Chimaraoke/bookings/internal/handlers"
	"github.com/Marvellous-Chimaraoke/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNum = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// to allow storage of reservaion details in session
	//gob.Register(models.Reservation{})
	// app.InProduction should be false in development and true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Firing up server on port %s", portNum))

	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
