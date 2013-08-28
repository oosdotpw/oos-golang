package handler

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"oos-go/lib"
	"oos-go/model"
)

type UserGetUser struct {
	lib.Handler
}

func GetGravatar(email string) string {
	md5Hash := md5.New()
	io.WriteString(md5Hash, email)
	md5String := hex.EncodeToString(md5Hash.Sum(nil))
	return "http://www.gravatar.com/avatar/" + md5String
}

func (h UserGetUser) Post() int {
	h.Init()

	if h.Filter("username", `^\w{3,16}$`, "bad_username") {
		return 200
	}

	username := h.PostValue["username"]

	m := model.GetAccount(username)

	dataJson := lib.Json{}
	dataJson["create_time"] = m.CreateTime.Unix()
	dataJson["avatar"] = GetGravatar(m.Email)
	dataJson["contact"] = m.Contact

	return h.Result(dataJson, false)
}
