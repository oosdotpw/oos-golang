package model

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"labix.org/v2/mgo/bson"
	"oos-go/db"
	"oos-go/utils"
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

func CheckAccount(username string, passwd string) bool {
	m := bson.M{"username": username, "passwd": toSHA256(passwd)}

	return db.Exist(Account, m)
}

func GetAccount(username string) AccountModel {
	m := AccountModel{}

	Account.Find(bson.M{"username": username}).One(&m)

	return m
}

func GetAccountByToken(token string) AccountModel {
	m := AccountModel{}

	findM := bson.M{"tokens.token": token, "tokens.alive": true}

	err := Account.Find(findM).One(&m)
	if err != nil {
		utils.Log(utils.ERR, err.Error())
	}

	return m
}

func InsertToken(userID bson.ObjectId, token string, ip string, ua string) {
	m := TokenModel{
		Token:       token,
		IP:          ip,
		UA:          ua,
		CreateTime:  time.Now(),
		ActiveTime:  time.Now(),
		ExpiredTime: time.Now().AddDate(0, 1, 0),
		Alive:       true,
	}

	Account.UpdateId(userID, bson.M{"$push": bson.M{"tokens": m}})
}

func CheckToken(token string) bool {
	m := AccountModel{}

	findM := bson.M{"tokens.token": token, "tokens.alive": true}
	selectM := bson.M{"tokens.$": 1}

	err := Account.Find(findM).Select(selectM).One(&m)
	if err != nil {
		return false
	}

	if m.Tokens[0].ExpiredTime.Before(time.Now()) {
		changeM := bson.M{"$set": bson.M{"tokens.$.alive": false}}

		err := Account.Update(findM, changeM)
		if err != nil {
			utils.Log(utils.ERR, err.Error())
		}
		return false
	}
	return true
}

func StopToken(token string) {
	findM := bson.M{"tokens.token": token}
	changeM := bson.M{"$set": bson.M{"tokens.$.Alive": false}}

	err := Account.Update(findM, changeM)
	if err != nil {
		utils.Log(utils.ERR, err.Error())
	}
}

func GetToken(token string) TokenModel {
	m := AccountModel{}

	findM := bson.M{"tokens.token": token, "tokens.alive": true}
	selectM := bson.M{"tokens.$": 1}

	err := Account.Find(findM).Select(selectM).One(&m)
	if err != nil {
		utils.Log(utils.INF, err.Error())
	}

	return m.Tokens[0]
}

func toSHA256(s string) string {
	h := sha256.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}
