package main

import (
	"log"
	"net/http"
)

func (a *application) homePageHandlers(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	a.session.Put(r.Context(), "remote_ip", remoteIP)

	err := renderTemplate(w, a.tc, &templateData{}, "home.page.gohtml")
	if err != nil {
		log.Fatal(err)
	}
}

func (a *application) aboutPageHandlers(w http.ResponseWriter, r *http.Request) {
	sMap := make(map[string]string)
	remote_ip := a.session.GetString(r.Context(), "remote_ip")
	sMap["remote_ip"] = remote_ip

	err := renderTemplate(w, a.tc, &templateData{SMap: sMap}, "about.page.gohtml")
	if err != nil {
		log.Fatal(err)
	}
}
