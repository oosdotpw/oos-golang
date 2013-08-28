package main

import (
	"net/http"
	"oos-go/handler"
	"oos-go/lib"
)

func main() {
	lib.Log(lib.INF, "oos start at "+lib.Config.Listen)

	maps := map[string]lib.HandlerInterface{
		"/account/signup":       handler.AccountSignup{},
		"/account/login":        handler.AccountLogin{},
		"/account/logout":       handler.AccountLogout{},
		"/account/session_info": handler.AccountSessionInfo{},
		"/user/get_user":        handler.UserGetUser{},
	}

	err := http.ListenAndServe(lib.Config.Listen, lib.NewHttpServer(maps))
	if err != nil {
		panic(err)
	}
}
