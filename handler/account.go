package handler

import (
	"math/rand"
	"oos-go/lib"
	"oos-go/model"
)

type AccountSignup struct {
	lib.Handler
}

type AccountLogin struct {
	lib.Handler
}

type AccountLogout struct {
	lib.Handler
}

type AccountSessionInfo struct {
	lib.Handler
}

func (h AccountSignup) Post() int {
	h.Init()

	if h.Filter("username", `^\w{3,16}$`, "bad_username") ||
		h.Filter("passwd", `^.{6,20}$`, "bad_passwd") ||
		h.Filter("email", `^\w+([-+.]\w+)*@\w+([-.]\w+)*$`, "bad_email") {
		return 200
	}

	username := h.PostValue["username"]
	email := h.PostValue["email"]
	passwd := h.PostValue["passwd"]
	contact := h.PostValue["contact"]

	if model.ExistAccount(username) {
		return h.Error("username_exits")
	}

	model.InsertAccount(username, passwd, email, contact)

	return h.Result(nil, false)
}

func (h AccountLogin) Post() int {
	h.Init()

	if h.Filter("username", `^\w{3,16}$`, "failure") ||
		h.Filter("passwd", `^.{6,20}$`, "failure") {
		return 200
	}

	username := h.PostValue["username"]
	passwd := h.PostValue["passwd"]
	ip := h.PostValue["IP"]
	ua := h.PostValue["UA"]

	if model.CheckAccount(username, passwd) == false {
		return h.Error("failure")
	}

	m := model.GetAccount(username)
	token := randomString(64)

	model.InsertToken(m.ObjectId, token, ip, ua)

	return h.Result(lib.Json{"token": token}, false)

}

func (h AccountLogout) Post() int {
	h.Init()

	if h.CheckToken() {
		return 200
	}

	model.StopToken(h.PostValue["token"])

	return h.Result(nil, false)
}

func (h AccountSessionInfo) Post() int {
	h.Init()

	if h.CheckToken() {
		return 200
	}

	m := model.GetToken(h.PostValue["token"])

	dataJson := lib.Json{
		"username": h.Account.Username,
		"expired":  m.ExpiredTime.Unix(),
	}

	return h.Result(dataJson, false)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}
