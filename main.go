package main

import (
	"net/http"
	"oos-go/handler"
	"oos-go/lib"
	"oos-go/utils"
)

func main() {
	utils.Log(utils.INF, "oos start at "+utils.Config.Listen)

	maps := map[string]lib.HandlerInterface{
		"/api/account/signup":          handler.AccountSignup{},
		"/api/account/login":           handler.AccountLogin{},
		"/api/account/logout":          handler.AccountLogout{},
		"/api/account/session_info":    handler.AccountSessionInfo{},
		"/api/user/get_user":           handler.UserGet{},
		"/api/post/new":                handler.PostNew{},
		"/api/post/get_post":           handler.PostGet{},
		"/api/post/reply":              handler.PostReply{},
		"/api/post/get_replies":        handler.PostGetReplys{},
		"/api/post/markup":             handler.PostMark{},
		"/api/post/fetch_by_number":    handler.FetchInit{},
		"/api/post/fetch_by_last_post": handler.FetchUpdate{},
		"/api/post/fetch_more":         handler.FetchMore{},
	}

	err := http.ListenAndServe(utils.Config.Listen, lib.NewHttpServer(maps))
	if err != nil {
		panic(err)
	}
}
