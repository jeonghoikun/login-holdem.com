package site

import (
	"strings"
	"time"
)

var Config *config

type Keywords []string

func (k *Keywords) String() string { return strings.Join(*k, ",") }

type config struct {
	Port          uint32
	Domain        string
	Author        string
	Title         string
	Description   string
	Keywords      *Keywords
	DatePublished time.Time
	DateModified  time.Time
	PhoneNumber   string
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func Init() {
	c := &config{}
	c.Port = uint32(8000)
	c.Domain = "domain.com"
	c.Author = "author"
	c.Title = "title"
	c.Description = "description"
	k := Keywords([]string{"keyword1", "keyword2", "keyword3", "keyword4.."})
	c.Keywords = &k
	c.DatePublished = date(2023, 8, 5)
	c.DateModified = date(2023, 8, 5)
	// 업종마다 전화번호가 다른경우 store/store.go 파일의 setPhoneNumber 함수에서 하드코딩
	c.PhoneNumber = "010-1234-1234"
	Config = c
}
