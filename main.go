package main

import (
	"handler"
	"lib"
	"net/http"
)

func main() {
	lib.Log(lib.INF, "oos start at "+lib.Config.Listen)

	maps := map[string]lib.HandlerInterface{
		"/account/signup": handler.AccountSignup{},
	}

	err := http.ListenAndServe(lib.Config.Listen, lib.NewHttpServer(maps))
	if err != nil {
		panic(err)
	}
}
