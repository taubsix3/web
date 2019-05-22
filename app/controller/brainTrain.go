package controller

import (
	"html/template"
	"net/http"
	"../model"
)

var tmpl *template.Template
// Is executed automatically on package load
func init() {
	tmpl = template.Must(template.ParseGlob("template/*.tmpl"))
}

// Index controller
func Index(w http.ResponseWriter, r *http.Request) {
	stats, _ := model.GetStats()
	tmpl.ExecuteTemplate(w, "index.tmpl", stats)
}

func Bearbeiten(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "bearbeiten.tmpl", nil)
}
func Bearbeiten2(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "bearbeiten2.tmpl", nil)
}

func Lernen(w http.ResponseWriter, r *http.Request) {
	_id := r.FormValue("_id")
	kasten, _ := model.GetKasten(_id)
	tmpl.ExecuteTemplate(w, "lernen.tmpl", kasten)
}

func Lernen2(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "lernen2.tmpl", nil)
}
func Anschauen(w http.ResponseWriter, r *http.Request) {
	_id := r.FormValue("_id")
	kasten, _ := model.GetKasten(_id)
	karten, _ := model.GetAllKarten()
	type Data struct {
		Kasten    model.Kasten
		AllKarten []map[string]interface{}
	}
	var data Data
	data.Kasten = kasten
	data.AllKarten = karten
	tmpl.ExecuteTemplate(w, "anschauen.tmpl", data)
}

func Karteikasten(w http.ResponseWriter, r *http.Request) {
	kasten, _ := model.GetAllKasten()
	tmpl.ExecuteTemplate(w, "karteikasten.tmpl", kasten)
}

func MeineKarteien(w http.ResponseWriter, r *http.Request) {
	kasten, _:= model.GetAllKasten()
	tmpl.ExecuteTemplate(w, "meinekarteien.tmpl",kasten)
}
func Profil(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "profil.tmpl", nil)
}
func Registrieren(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "registrieren.tmpl", nil)
}
