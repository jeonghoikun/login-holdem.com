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
	PhoneNumber   map[string]string
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
	c.PhoneNumber = map[string]string{
		"Global": "010-1234-1234",
		// 특정 업종에 따라 폰번호가 달라야하는경우 아래에 하드코딩으로 추가
		// "Hobba": "010-2345-2345,
		// "Club":  "010-4321-4321,
	}
	Config = c
}
