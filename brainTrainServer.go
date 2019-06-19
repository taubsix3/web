package main

import (
	"net/http"
	"../BrainTrain.1/app/controller"
)

func main() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/registrieren", controller.Registrieren)
	http.HandleFunc("/anschauen", controller.Anschauen)
	http.HandleFunc("/bearbeiten", controller.Bearbeiten)
	http.HandleFunc("/bearbeiten2", controller.Bearbeiten2)
	http.HandleFunc("/index", controller.Index)
	http.HandleFunc("/karteikasten", controller.Karteikasten)
	http.HandleFunc("/lernen", controller.Lernen)
	http.HandleFunc("/lernen2", controller.Lernen2)
	http.HandleFunc("/meinekarteien", controller.MeineKarteien)
	http.HandleFunc("/profil", controller.Profil)

	http.HandleFunc("/register", controller.Register)
http.HandleFunc("/add-user", controller.AddUser)
http.HandleFunc("/login", controller.Login)
http.HandleFunc("/authenticate-user", controller.AuthenticateUser)
http.HandleFunc("/logout", controller.Logout)


	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/css/", http.StripPrefix("/", fs))
	http.Handle("/fonts/", http.StripPrefix("/", fs))
	http.Handle("/js/", http.StripPrefix("/", fs))
	http.Handle("/img/", http.StripPrefix("/", fs))


	server := http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()
}
