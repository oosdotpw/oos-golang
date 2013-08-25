package lib

import (
	"log"
)

const (
	ERR = iota
	INF
	WAR
	DEG
)

func Log(level int, info interface{}) {
	if level <= Config.Loglevel {
		log.Println(info)
	}
}
