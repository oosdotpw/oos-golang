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
		"/account/signup":       handler.AccountSignup{},
		"/account/login":        handler.AccountLogin{},
		"/account/logout":       handler.AccountLogout{},
		"/account/session_info": handler.AccountSessionInfo{},
		"/user/get_user":        handler.UserGetUser{},
	}

	err := http.ListenAndServe(utils.Config.Listen, lib.NewHttpServer(maps))
	if err != nil {
		panic(err)
	}
}
