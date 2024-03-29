package model

import (
	couchdb "github.com/leesper/couchdb-golang"
	)



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
Nummer string `json:"nummer"`
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

//GetIdKarten from DB with ID
func GetIdKarten(k Kasten) ([]map[string]interface{}, error) {
	id := k.ID
	allKarten, err := btDB.QueryJSON(`
	{
		"selector": {
			 "$and": [
				 {"type":{"$eq": "karten"}},
				 {"kasten":{"$eq": "` + id + `"}}
			 ]

		}
	 }`)
	if err != nil {
		return nil, err
	} else {

		return allKarten, nil
	}
}
