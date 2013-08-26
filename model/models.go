package model

import (
	"labix.org/v2/mgo"
	"oos-go/db"
)

var Account *mgo.Collection

func init() {
	Account = db.GetCollection("account")
}
