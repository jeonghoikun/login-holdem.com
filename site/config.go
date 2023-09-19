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
	c.Port = uint32(4102)
	c.Domain = "holdem-kingdom.com"
	c.Author = "이사장"
	c.Title = "홀덤 킹덤"
	c.Description = "홀덤 킹덤은 강남과 압구정 지역에서 홀덤 포커를 즐길 수 있는 탁월한 공간으로, 카드 게임 팬들에게는 꼭 방문해야 하는 장소 중 하나입니다. 이 독특한 홀덤 플레이 지역은 그야말로 홀덤 판의 왕국으로 불립니다. 강남홀덤과 압구정홀덤 두 곳에서 홀덤 판을 즐길 수 있으며, 그 두 곳은 서로 다른 분위기와 스타일을 가지고 있어 모든 유형의 홀덤 팬들에게 맞춤형 경험을 제공합니다."
	k := Keywords([]string{"강남홀덤", "압구정홀덤", "홀덤킹덤"})
	c.Keywords = &k
	c.DatePublished = date(2023, 9, 19)
	c.DateModified = date(2023, 9, 19)
	// 업종마다 전화번호가 다른경우 store/store.go 파일의 setPhoneNumber 함수에서 하드코딩
	c.PhoneNumber = "010-5900-9210"
	c.SearchEngineConnection = &searchEngineConnection{
		Google: "uB7ecmImPZewv2LwvJ68mQiRh31D14Tjdv2kuqqPiho",
	}
	Config = c
}
