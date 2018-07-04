package api

import (
	"github.com/mitsukomegumi/Crypto-Go/src/accounts"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SetupAccountRoutes - setup necessary routes for accout database
func SetupAccountRoutes(db *mgo.Database) error {
	go func() error {
		pErr := setPosts(db)

		if pErr != nil {
			return pErr
		}
		return nil
	}()

	err := setGets(db)

	if err != nil {
		return err
	}

	return nil
}

func setGets(db *mgo.Database) error {
	req, err := NewRequestServer(":username", "/api/accounts", "GET", db, db, "username")

	if err != nil {
		return err
	}

	err = req.AttemptToServeRequests()

	if err != nil {
		return err
	}

	return nil
}

func setPosts(db *mgo.Database) error {
	postReq, rErr := NewRequestServer("POST", "/api/accounts", "POST", nil, db, "/:username/:email/:pass")

	if rErr != nil {
		return rErr
	}

	pErr := postReq.AttemptToServeRequests()

	if pErr != nil {
		panic(rErr)
	}

	return nil
}

func addAccount(database *mgo.Database, account *accounts.Account) error {
	c := database.C("accounts")

	err := c.Insert(account)

	if err != nil {
		return err
	}

	return nil
}

func findAccount(database *mgo.Database, username string) (*accounts.Account, error) {
	c := database.C("accounts")

	result := accounts.Account{}

	err := c.Find(bson.M{"username": username}).One(&result)
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func findValue(database *mgo.Database, collection string, key string, value string) (interface{}, error) {
	c := database.C(collection)

	result := make(map[string]interface{})

	err := c.Find(bson.M{key: value}).One(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
