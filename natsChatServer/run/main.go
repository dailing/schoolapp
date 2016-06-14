package main

import (
	"git.oschina.net/dddailing/schoolapp/natsChatServer"
	"github.com/dailing/levlog"
	"time"
)

func main() {
	levlog.Info("good")
	server := natsChatServer.NewServer()
	server.Start()
	for {
		time.Sleep(time.Second * 10)
	}
}
