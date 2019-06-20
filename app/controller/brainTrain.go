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
	karteikarten, _ := model.GetAllKarten()
		nutzer, _ := model.GetAllUser()
		karteien, _ := model.GetAllKasten()
		anzahlKarteikarten := len(karteikarten)
		anzahlNutzer := len(nutzer)
		anzahlKarteien := len(karteien)
		data := struct {
			//Karteikarten       *[]map[string]interface{}
			Karteikartenanzahl int
			Nutzeranzahl       int
			Karteienanzahl     int
		}{
			//&karteikarten,
			anzahlKarteikarten,
			anzahlNutzer,
			anzahlKarteien,
		}
		tmpl.ExecuteTemplate(w, "index.tmpl", data)
	}

	func Home(w http.ResponseWriter, r *http.Request) {

			tmpl.ExecuteTemplate(w, "home.tmpl", nil)
		}

// Bearbeiten controller
func Bearbeiten(w http.ResponseWriter, r *http.Request) {
	_id := r.FormValue("_id")
	kasten, _ := model.GetKasten(_id)

	tmpl.ExecuteTemplate(w, "bearbeiten.tmpl", kasten)
}

//Bearbeiten2 controller
func Bearbeiten2(w http.ResponseWriter, r *http.Request) {
	_id := r.FormValue("_id")
	kasten, _ := model.GetKasten(_id)
	karten, _ := model.GetIdKarten(kasten)

	type Data struct {
		Kasten    model.Kasten
		AllKarten []map[string]interface{}
	}
	var data Data
	data.Kasten = kasten
	data.AllKarten = karten

	tmpl.ExecuteTemplate(w, "bearbeiten2.tmpl", data)
}

//Karteianschauen controller
func Anschauen(w http.ResponseWriter, r *http.Request) {
	_id := r.FormValue("_id")
	kasten, _ := model.GetKasten(_id)
	karten, _ := model.GetIdKarten(kasten)

	type Data struct {
		Kasten    model.Kasten
		AllKarten []map[string]interface{}
	}
	var data Data
	data.Kasten = kasten
	data.AllKarten = karten

	tmpl.ExecuteTemplate(w, "anschauen.tmpl", data)
}

//Karteikasten controller
func Karteikasten(w http.ResponseWriter, r *http.Request) {
	kasten, _ := model.GetAllKasten()

	tmpl.ExecuteTemplate(w, "karteikasten.tmpl", kasten)
}

//Lernen controller
func Lernen(w http.ResponseWriter, r *http.Request) {
_id:= r.FormValue("_id")
kasten, _ :=model.GetKasten(_id)
karten, _ := model.GetIdKarten(kasten)

type Data struct {
	Kasten    model.Kasten
	AllKarten []map[string]interface{}
}
var data Data
data.Kasten = kasten
data.AllKarten = karten

	tmpl.ExecuteTemplate(w, "lernen.tmpl", data)
}

func Lernen2(w http.ResponseWriter, r *http.Request) {
_id:= r.FormValue("_id")
kasten, _ :=model.GetKasten(_id)
karten, _ := model.GetIdKarten(kasten)

type Data struct {
	Kasten    model.Kasten
	AllKarten []map[string]interface{}
}
var data Data
data.Kasten = kasten
data.AllKarten = karten

	tmpl.ExecuteTemplate(w, "lernen2.tmpl", data)
}


//Meinekarteikasten controller
func MeineKarteien(w http.ResponseWriter, r *http.Request) {
	kasten, _ := model.GetAllKasten()

	tmpl.ExecuteTemplate(w, "meinekarteien.tmpl", kasten)
}

//Profil controller
func Profil(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "profil.tmpl", nil)
}

//Registrieren controller
func Registrieren(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "registrieren.tmpl", nil)
}
