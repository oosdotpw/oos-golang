package handler

import (
	"oos-go/lib"
	"oos-go/model"
	"strconv"
)

type PostNew struct {
	lib.Handler
}

type PostGet struct {
	lib.Handler
}

type PostReply struct {
	lib.Handler
}

type PostGetReplys struct {
	lib.Handler
}

type PostMark struct {
	lib.Handler
}

type FetchInit struct {
	lib.Handler
}

type FetchUpdate struct {
	lib.Handler
}

type FetchMore struct {
	lib.Handler
}

func (h PostNew) Post() int {
	h.Init()

	if h.CheckToken() {
		return 200
	}

	if h.Filter("content", `.+`, "bad_content") {
		return 200
	}

	content := h.PostValue["content"]

	id := model.InsertPost(h.Account.ObjectId, content)

	return h.Result(lib.Json{"id": id}, false)
}

func (h PostGet) Post() int {
	h.Init()

	id := h.PostValue["id"]

	if !model.CheckObjectID(id) || !model.ExistPost(id) {
		return h.Error("bad_reply_post")
	}

	postM := model.GetPost(id)
	userM := model.GetAccountByID(postM.UserID)

	dataJson := lib.Json{
		"content": postM.Content,
		"time":    postM.CreateTime,
		"replys":  len(postM.Replys),
		"author": lib.Json{
			"username": userM.Username,
			"avatar":   GetGravatar(userM.Email),
		},
	}

	return h.Result(dataJson, false)
}

func (h PostReply) Post() int {
	h.Init()

	if h.CheckToken() {
		return 200
	}

	if h.Filter("content", `.+`, "bad_content") {
		return 200
	}

	content := h.PostValue["content"]
	postID := h.PostValue["reply_post"]

	if !model.CheckObjectID(postID) || !model.ExistPost(postID) {
		return h.Error("bad_reply_post")
	}

	objID := model.InsertReply(h.Account.ObjectId, postID, content)

	return h.Result(lib.Json{"id": objID.Hex()}, false)
}

func (h PostGetReplys) Post() int {
	h.Init()

	id := h.PostValue["id"]

	if !model.CheckObjectID(id) || !model.ExistPost(id) {
		return h.Error("bad_reply_post")
	}

	replysM := model.GetReplys(id)

	resultJson := make([]map[string]interface{}, 0, 100)

	for _, replyM := range replysM {
		userM := model.GetAccountByID(replyM.UserID)

		replyJson := lib.Json{
			"content": replyM.Content,
			"time":    replyM.CreateTime,
			"author": lib.Json{
				"username": userM.Username,
				"avatar":   GetGravatar(userM.Email),
			},
		}

		resultJson = append(resultJson, replyJson)
	}

	return h.Result(lib.Json{"result": resultJson}, false)
}

func (h PostMark) Post() int {
	h.Init()

	if h.CheckToken() {
		return 200
	}

	postID := h.PostValue["id"]
	markType := h.PostValue["type"]

	if !model.CheckObjectID(postID) || !model.ExistPost(postID) {
		return h.Error("bad_reply_post")
	}

	if !checkMarkType(markType) {
		return h.Error("bad_type")
	}

	model.InsertMark(h.Account.ObjectId, postID, markType)

	return h.Result(nil, false)

}

func (h FetchInit) Post() int {
	h.Init()

	if h.CheckToken() {
		return 200
	}

	num, err := strconv.Atoi(h.PostValue["num"])
	if err != nil {
		return h.Error("bad_num")
	}

	postsM := model.FetchNewest(num)

	resultJson := make([]string, len(postsM), len(postsM))
	for i, postM := range postsM {
		resultJson[i] = postM.ObjectId.Hex()
	}

	return h.Result(lib.Json{"result": resultJson}, false)
}

func (h FetchUpdate) Post() int {
	h.Init()

	if h.CheckToken() {
		return 200
	}

	id := h.PostValue["id"]
	if !model.CheckObjectID(id) || !model.ExistPost(id) {
		return h.Error("bad_reply_post")
	}

	postsM := model.FetchNewer(id, 100)

	resultJson := make([]string, len(postsM), len(postsM))
	for i, postM := range postsM {
		resultJson[i] = postM.ObjectId.Hex()
	}

	return h.Result(lib.Json{"result": resultJson}, false)
}

func (h FetchMore) Post() int {
	h.Init()

	if h.CheckToken() {
		return 200
	}

	id := h.PostValue["id"]
	if !model.CheckObjectID(id) || !model.ExistPost(id) {
		return h.Error("bad_reply_post")
	}

	postsM := model.FetchOlder(id, 100)

	resultJson := make([]string, len(postsM), len(postsM))
	for i, postM := range postsM {
		resultJson[i] = postM.ObjectId.Hex()
	}

	return h.Result(lib.Json{"result": resultJson}, false)
}

func checkMarkType(mark string) bool {
	markTypes := []string{"like", "dislike", "spam"}
	for _, v := range markTypes {
		if v == mark {
			return true
		}
	}
	return false
}
