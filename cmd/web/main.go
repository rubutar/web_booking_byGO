package main

import (
	"fmt"
	"github.com/rubutar/web_booking_byGO/pkg/config"
	"github.com/rubutar/web_booking_byGO/pkg/handler"
	"github.com/rubutar/web_booking_byGO/pkg/render"
	"log"
	"net/http"
)

const portNumb = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumb))
	_ = http.ListenAndServe(portNumb, nil)
}
