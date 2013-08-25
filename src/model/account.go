package model

import (
	"crypto/sha256"
	"db"
	"encoding/hex"
	"io"
	"labix.org/v2/mgo/bson"
	"time"
)

type TokenModel struct {
	Token       string    `token`
	IP          string    `ip`
	UA          string    `ua`
	CreateTime  time.Time `create_time`
	ActiveTime  time.Time `last_active`
	ExpiredTime time.Time `expired`
	Alive       bool      `alive`
}

type AccountModel struct {
	ObjectId   bson.ObjectId `_id,omitempty`
	Username   string        `username`
	Passwd     string        `passwd`
	Contact    string        `contact`
	CreateTime time.Time     `create_time`
	Email      string        `email`
	Tokens     []TokenModel  `tokens,omitempty`
}

func InsertAccount(username string, passwd string, email string, contact string) {
	toSHA256 := func(s string) string {
		h := sha256.New()
		io.WriteString(h, s)
		return hex.EncodeToString(h.Sum(nil))
	}

	m := AccountModel{
		Username:   username,
		Passwd:     toSHA256(passwd),
		Contact:    contact,
		CreateTime: time.Now(),
		Email:      email,
		Tokens:     []TokenModel{},
	}

	Account.Insert(m)
}

func ExistAccount(username string) bool {
	return db.Exist(Account, bson.M{"username": username})
}
