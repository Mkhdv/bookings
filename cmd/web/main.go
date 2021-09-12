package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/mkhdv/bookings/pkg/config"
	"github.com/mkhdv/bookings/pkg/handlers"
	"github.com/mkhdv/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

//Home page handler
//func Home (w http.ResponseWriter, r *http.Request) {
//	renderTemplate(w, "testhome.page.tmpl")
//}

// About page handler
//func About (w http.ResponseWriter, r *http.Request) {
//	renderTemplate(w, "testabout.page.tmpl")
//}

//func renderTemplate(w http.ResponseWriter, tmpl string) {
//	parsedTemplate, _ := template.ParseFiles("../../templates/" + tmpl)
//	err := parsedTemplate.Execute(w, nil)
//	if err != nil {
//		fmt.Println("error", err)
//		return
//	}
//}

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager


func main() {

	//change this to true when in production
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


	app.UseCache = false
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}


