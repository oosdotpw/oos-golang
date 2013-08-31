package model

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"oos-go/db"
)

var Account *mgo.Collection
var Post *mgo.Collection
var Mark *mgo.Collection

func init() {
	Account = db.GetCollection("account")
	Post = db.GetCollection("post")
	Mark = db.GetCollection("mark")
}

func CheckObjectID(ID string) bool {
	return bson.IsObjectIdHex(ID)
}
