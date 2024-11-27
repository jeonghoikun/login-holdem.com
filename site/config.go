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
	c.Port = uint32(4103)
	c.Domain = "login-holdem.com"
	c.Author = "김딜러"
	c.Title = "강남홀덤 압구정홀덤 로그인"
	c.Description = "당신을 강남홀덤 압구정홀덤 로그인의 세계로 초대합니다. 여기는 강남홀덤과 압구정홀덤이 한 곳에서 만나는 뜨거운 경쟁의 무대입니다. 강남홀덤 압구정홀덤 로그인에서는 게임의 짜릿함을 느낄 수 있으며, 다양한 레벨의 플레이어와 함께 경험을 공유할 수 있습니다."
	k := Keywords([]string{"강남홀덤", "압구정홀덤", "로그인 홀덤", "강남홀덤 압구정홀덤 로그인", "로그인"})
	c.Keywords = &k
	c.DatePublished = date(2023, 10, 16)
	c.DateModified = date(2024, 1, 10)
	// 업종마다 전화번호가 다른경우 store/store.go 파일의 setPhoneNumber 함수에서 하드코딩
	c.PhoneNumber = "010-4636-3309"
	c.SearchEngineConnection = &searchEngineConnection{
		Google: "e9dafNYxmGa1DwOGWBXQ5T213mPPAxIBc0Urpn0qJMU",
	}
	Config = c
}
