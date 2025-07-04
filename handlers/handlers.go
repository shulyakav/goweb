package handlers

import (
	"net/http"
	"shulyakav/goweb/renders"
)

func Home(w http.ResponseWriter, r *http.Request) {

	renders.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {

	renders.RenderTemplate(w, "about.page.tmpl")
}
