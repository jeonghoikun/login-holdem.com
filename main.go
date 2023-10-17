package main

import (
	"log"
	"time"

	"github.com/jeonghoikun/login-holdem.com/server"
	"github.com/jeonghoikun/login-holdem.com/site"
)

func init() {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	time.Local = loc
	site.Init()
}

func main() {
	s := server.New(site.Config.Port)
	log.Fatal(s.Run())
}
