package model

import (
	"encoding/json"

	couchdb "../../couchdb-golang-master"
	)

//User data structure
type User struct {
	ID       string `json:"_id"`
	Rev      string `json:"_rev"`
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
	Status   bool   `json:"status"`
	Kasten   []Kasten

	*couchdb.Document
}

//Kasten data structure
type Kasten struct {
	ID             string `json:"_id"`
	Rev            string `json:"_rev"`
	Type           string `json:"type"`
	Unterkategorie string `json:"unterkategorie"`
	Oberkategorie  string `json:"oberkategorie"`
	Titel          string `json:"titel"`
	Beschreibung   string `json:"beschreibung"`
	Privat         bool   `json:"privat"`
	Fortschritt    int    `json:"fortschritt"`
	Ersteller      string `json:"ersteller"`

	*couchdb.Document
}

//Karten data structure
type Karten struct {
	ID      string `json:"_id"`
	Rev     string `json:"_rev"`
	Type    string `json:"type"`
	Titel   string `json:"titel"`
	Frage   string `json:"frage"`
	Antwort string `json:"antwort"`
	Kasten  string `json:"kasten"`

	*couchdb.Document
}

//Statistik data structure
type Statistik struct {
	ID            string `json:"_id"`
	Rev           string `json:"_rev"`
	Type          string `json:"type"`
	CounterUser   int    `json:"counterUser"`
	CounterKarten int    `json:"counterKarten"`
	CounterKasten int    `json:"counterKasten"`

	*couchdb.Document
}

// Db handle
var btDB *couchdb.Database

func init() {
	var err error
	btDB, err = couchdb.NewDatabase("http://localhost:5984/braintrain")
	if err != nil {
		panic(err)
	}
}

//GetAllUser from DB
func GetAllUser() ([]map[string]interface{}, error) {
	allUser, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "user"
			 }
		}
	 }`)
	 
	if err != nil {
		return nil, err
	} else {
		return allUser, nil
	}
}

//GetStats from DB
func GetStats() ([]map[string]interface{}, error) {
	allStats, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "statistik"
			 }
		}
	 }`)
	if err != nil {
		return nil, err
	} else {
		return allStats, nil
	}
}

//GetAllKarten from DB
func GetAllKarten() ([]map[string]interface{}, error) {
	allKarten, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "karten"

			 }

		}
	 }`)
	if err != nil {
		return nil, err
	} else {

		return allKarten, nil
	}
}

//GetAllKasten from DB
func GetAllKasten() ([]map[string]interface{}, error) {
	allKasten, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "kasten"

			 }

		}
	 }`)
	if err != nil {
		return nil, err
	} else {

		return allKasten, nil
	}
}

// GetUser with the provided id from DB
func GetUser(id string) (User, error) {
	u, err := btDB.Get(id, nil)
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:       u["_id"].(string),
		Rev:      u["_rev"].(string),
		Type:     u["type"].(string),
		Username: u["username"].(string),
		Password: u["password"].(string),
		Mail:     u["mail"].(string),
		Status:   u["status"].(bool),
	}
	return user, nil
}

// GetKasten with the provided id from DB
func GetKasten(id string) (Kasten, error) {
	k, err := btDB.Get(id, nil)
	if err != nil {
		return Kasten{}, err
	}

	kasten := Kasten{
		ID:             k["_id"].(string),
		Rev:            k["_rev"].(string),
		Type:           k["type"].(string),
		Unterkategorie: k["unterkategorie"].(string),
		Oberkategorie:  k["oberkategorie"].(string),
		Titel:          k["titel"].(string),
		Beschreibung:   k["beschreibung"].(string),
		Privat:         k["privat"].(bool),
		Fortschritt:    int(k["fortschritt"].(float64)),
		Ersteller:      k["ersteller"].(string)}
	return kasten, nil

}

// ---------------------------------------------------------------------------
// Internal helper functions
// ---------------------------------------------------------------------------

// Convert from struct to map[string]interface{} as required by Set() method
func user2Map(u User) map[string]interface{} {
	var doc map[string]interface{}
	tJSON, _ := json.Marshal(u)
	json.Unmarshal(tJSON, &doc)

	return doc
}
