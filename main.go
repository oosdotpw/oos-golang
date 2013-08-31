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
		"/account/signup":          handler.AccountSignup{},
		"/account/login":           handler.AccountLogin{},
		"/account/logout":          handler.AccountLogout{},
		"/account/session_info":    handler.AccountSessionInfo{},
		"/user/get_user":           handler.UserGet{},
		"/post/new":                handler.PostNew{},
		"/post/get_post":           handler.PostGet{},
		"/post/reply":              handler.PostReply{},
		"/post/get_replies":        handler.PostGetReplys{},
		"/post/markup":             handler.PostMark{},
		"/post/fetch_by_number":    handler.FetchInit{},
		"/post/fetch_by_last_post": handler.FetchUpdate{},
		"/post/fetch_more":         handler.FetchMore{},
	}

	err := http.ListenAndServe(utils.Config.Listen, lib.NewHttpServer(maps))
	if err != nil {
		panic(err)
	}
}
