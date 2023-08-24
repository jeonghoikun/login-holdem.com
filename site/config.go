package site

import (
	"strings"
	"time"
)

var Config *config

type Keywords []string

func (k *Keywords) String() string { return strings.Join(*k, ",") }

type searchEngineConnection struct {
	Google string
}

type config struct {
	Port                   uint32
	Domain                 string
	Author                 string
	Title                  string
	Description            string
	Keywords               *Keywords
	DatePublished          time.Time
	DateModified           time.Time
	PhoneNumber            string
	SearchEngineConnection *searchEngineConnection
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func Init() {
	c := &config{}
	c.Port = uint32(4101)
	c.Domain = "gn-ag-holdem.com"
	c.Author = "kevin"
	c.Title = "강남홀덤 압구정홀덤 게이트"
	c.Description = "강남홀덤 압구정홀덤 게이트는 즐거운 홀덤 포커 경험을 제공하는 최상의 오프홀덤 입니다. 강남홀덤 압구정홀덤! 프로페셔널 딜러, 다양한 테이블, 친근한 분위기로 진정한 홀덤 열정을 만끽하세요. 초보자부터 베테랑까지 모두에게 적합한 플레이 환경을 제공합니다."
	k := Keywords([]string{"강남홀덤", "압구정홀덤", "게이트홀덤"})
	c.Keywords = &k
	c.DatePublished = date(2023, 8, 23)
	c.DateModified = date(2023, 8, 23)
	// 업종마다 전화번호가 다른경우 store/store.go 파일의 setPhoneNumber 함수에서 하드코딩
	c.PhoneNumber = "010-5900-9210"
	c.SearchEngineConnection = &searchEngineConnection{
		Google: "uB7ecmImPZewv2LwvJ68mQiRh31D14Tjdv2kuqqPiho",
	}
	Config = c
}
