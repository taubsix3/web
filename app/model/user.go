package model

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)



// Add User
func (user User) Add() (err error) {
	// Check wether username already exists
	userInDB, err := GetUserByUsername(user.Username)
	if err == nil && userInDB.Username == user.Username {
		return errors.New("username exists already")
	}

	// Hash password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)

	user.Password = b64HashedPwd
	user.Type = "user"

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
