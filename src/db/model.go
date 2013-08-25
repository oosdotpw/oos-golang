package db

import (
	"labix.org/v2/mgo"
	"lib"
)

var (
	mongo  *mgo.Session
	dbname string
)

func init() {
	var err error
	mongo, err = mgo.Dial(lib.Config.Mongodb)
	if err != nil {
		panic(err)
	}

	dbname = lib.Config.Dbname
}

func Exist(c *mgo.Collection, query interface{}) bool {
	n, _ := c.Find(query).Count()
	if n > 0 {
		return true
	}
	return false
}

func GetCollection(name string) *mgo.Collection {
	return mongo.DB(dbname).C(name)
}
