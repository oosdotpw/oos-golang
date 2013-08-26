package handler

import (
	"oos-go/lib"
	"oos-go/model"
)

type AccountSignup struct {
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
