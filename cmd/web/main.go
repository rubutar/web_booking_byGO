package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
	"web_booking_byGO/pkg/config"
	"web_booking_byGO/pkg/handler"
	"web_booking_byGO/pkg/render"
)

const portNumb = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//change this to true when in production mode
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumb))

	srv := &http.Server{
		Addr:    portNumb,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
