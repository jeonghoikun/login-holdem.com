package main

import (
	"log"
	"time"

	"github.com/jeonghoikun/gn-ag-holdem.com/server"
	"github.com/jeonghoikun/gn-ag-holdem.com/site"
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
