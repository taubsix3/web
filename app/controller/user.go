package controller

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"../model"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store *sessions.CookieStore

func init() {
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key := make([]byte, 32)
	rand.Read(key)
	store = sessions.NewCookieStore(key)
}

// Register controller
func Register(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "registrieren.tmpl", nil)
}

// AddUser controller
func AddUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	mail := r.FormValue("mail")


	user := model.User{}
	user.Username = username
	user.Password = password
	user.Mail= mail

	err := user.Add()
	if err != nil {
		data := struct {
			ErrorMsg string
		}{
			ErrorMsg: "Username already exists!",
		}
		tmpl.ExecuteTemplate(w, "index.tmpl", data)
	} else {
		tmpl.ExecuteTemplate(w, "index.tmpl", nil)
	}
}

// Login controller
func Login(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.tmpl", nil)
}

// AuthenticateUser controller
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var user = model.User{}
	var data = struct {
		ErrorMsg string
	}{
		ErrorMsg: "Username and/or password wrong!",
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Authentication
	user, err = model.GetUserByUsername(username)

	if err == nil {
		// decode base64 String to []byte
		passwordDB, _ := base64.StdEncoding.DecodeString(user.Password)
		err = bcrypt.CompareHashAndPassword(passwordDB, []byte(password))
		if err == nil {
			session, _ := store.Get(r, "session")

			// Set user as authenticated
			session.Values["authenticated"] = true
			session.Values["username"] = username
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			tmpl.ExecuteTemplate(w, "index.tmpl", data)
		}
	} else {
		tmpl.ExecuteTemplate(w, "index.tmpl", data)
	}
}

// Logout controller
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Save(r, w)

	tmpl.ExecuteTemplate(w, "index.tmpl", nil)
}

// Auth is an authentication handler
func Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			h(w, r)
		}
	}
}
