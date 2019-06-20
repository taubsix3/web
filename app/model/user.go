package model

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	couchdb "github.com/leesper/couchdb-golang"
	"time"

	"golang.org/x/crypto/bcrypt"
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
	Created   time.Time   `json:"created"`
	Kasten   []Kasten

	*couchdb.Document
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

// Add User
func (user User) Add() (err error) {
	// Check wether username already exists
	userInDB, err := GetUserByUsername(user.Username)
	//mailInDB, err := GetUserByMail(user.Mail)

	if err == nil && userInDB.Username == user.Username {
		return errors.New("username already exists")
	}
//	if err == nil && mailInDB.Mail == user.Mail {
	//	return errors.New("mail already exists")
	//}
	// Hash password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)

	user.Password = b64HashedPwd
	user.Type = "user"
user.Created = time.Now()

	// Convert Todo struct to map[string]interface as required by Save() method
	u, err := user2Mapp(user)

	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(u, "_id")
	delete(u, "_rev")

	// Add todo to DB
	_, _, err = btDB.Save(u, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return err
}

// GetUserByUsername retrieve User by username
func GetUserByUsername(username string) (user User, err error) {
	if username == "" {
		return User{}, errors.New("no username provided")
	}

	query := `
	{
		"selector": {
			 "type": "user",
			 "username": "%s"
		}
	}`
	u, err := btDB.QueryJSON(fmt.Sprintf(query, username))
	if err != nil || len(u) != 1 {
		return User{}, err
	}

	user, err = map2User(u[0])
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUserByMail(mail string) (user User, err error) {
	if mail == "" {
		return User{}, errors.New("no mailadress provided")
	}

	query := `
	{
		"selector": {
			 "type": "user",
			 "mail": "%s"
		}
	}`
	u, err := btDB.QueryJSON(fmt.Sprintf(query, mail))
	if err != nil || len(u) != 1 {
		return User{}, err
	}

	user, err = map2User(u[0])
	if err != nil {
		return User{}, err
	}

	return user, nil
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
		Created:  u["created"].(time.Time),
	}
	return user, nil
}

// ---------------------------------------------------------------------------
// Internal helper functions
// ---------------------------------------------------------------------------

// Convert from User struct to map[string]interface{} as required by golang-couchdb methods
func user2Mapp(u User) (user map[string]interface{}, err error) {
	uJSON, err := json.Marshal(u)
	json.Unmarshal(uJSON, &user)

	return user, err
}

// Convert from map[string]interface{} to User struct as required by golang-couchdb methods
func map2User(user map[string]interface{}) (u User, err error) {
	uJSON, err := json.Marshal(user)
	json.Unmarshal(uJSON, &u)

	return u, err
}
