package main

import (
	"handler"
	"lib"
	"net/http"
)

func main() {
	lib.LoadConfig("./config.yaml")
	lib.Log(lib.INF, "oos start at "+lib.Config.Listen)

	maps := map[string]lib.HandlerInterface{
		"/account/signup": handler.AccountSignup{},
	}

	err := http.ListenAndServe(lib.Config.Listen, lib.NewHttpHandler(maps))
	if err != nil {
		panic(err)
	}
}
