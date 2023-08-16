package main

import (
	"log"
	"time"

	"github.com/jeonghoikun/webserver/server"
	"github.com/jeonghoikun/webserver/site"
	"github.com/jeonghoikun/webserver/store"
)

func init() {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	time.Local = loc
	site.Init()
	if err := store.Init(); err != nil {
		panic(err)
	}
}

func main() {
	s := server.New(site.Config.Port)
	log.Fatal(s.Run())
}
