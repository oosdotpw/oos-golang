package model

import (
	"db"
	"labix.org/v2/mgo"
)

var Account *mgo.Collection

func init() {
	Account = db.GetCollection("account")
}
