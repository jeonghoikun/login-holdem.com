package store

import (
	"strings"
)

var stores []*Store = []*Store{}

func Get(do, si, dong, storeType, title string) (o *Store, has bool) {
	for _, s := range stores {
		if s.Location.Do == do && s.Location.Si == si && s.Location.Dong == dong &&
			s.Type == storeType && s.Title == title {
			return s, true
		}
	}
	return nil, false
}

func ListAllStores() []*Store { return stores }

func ListStoresByDo(do string) []*Store {
	list := []*Store{}
	for _, s := range stores {
		if s.Location.Do == do {
			list = append(list, s)
		}
	}
	return list
}

func ListStoresByDoSi(do, si string) []*Store {
	list := []*Store{}
	for _, s := range stores {
		if s.Location.Do == do && s.Location.Si == si {
			list = append(list, s)
		}
	}
	return list
}

func ListStoresByDoSiDong(do, si, dong string) []*Store {
	list := []*Store{}
	for _, s := range stores {
		if s.Location.Do == do && s.Location.Si == si && s.Location.Dong == dong {
			list = append(list, s)
		}
	}
	return list
}

type Location struct {
	// Do: 서울
	Do string
	// Si: 강남구
	Si string
	// Dong: 역삼동
	Dong string
	// Address: 822-5
	Address string
	// GoogleMapSrc: iframe google map의 src속성 값
	GoogleMapSrc string
}

type Keywords []string

func (k *Keywords) String() string { return strings.Join(*k, ",") }

type Active struct {
	// IsPermanentClosed: 영업중=true 폐업=false
	IsPermanentClosed bool
	// Reason: 폐업상태일 경우에만 입력
	Reason string
}

type Store struct {
	Location *Location
	// Type: 업종 하드코딩
	Type string
	// Title: 가게이름 하드코딩
	Title string
	// Description: 가게 설명 하드코딩
	Description string
	// Keywords: 하드코딩 X. 서버 시작시 지역명, 가게이름, 업종 등으로 자동 초기화 됨
	Keywords string
	// Active: 영업, 폐업 유무와 폐업사유 하드코딩
	Active *Active
	// PhoneNumber: 하드코딩 X.
	PhoneNumber string
}

func initHighPublic() {
	storeType := "하이퍼블릭"
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "822-5",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.3849794120856!2d127.02926860000001!3d37.4988373!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca159d7d08f47%3A0x19ac7457d361928!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDthYztl6TrnoDroZwgMTEx!5e0!3m2!1sko!2skr!4v1661153125692!5m2!1sko!2skr",
		},
		Type:        storeType,
		Title:       "메이커",
		Description: "메이커 is ...",
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "822-5",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.3849794120856!2d127.02926860000001!3d37.4988373!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca159d7d08f47%3A0x19ac7457d361928!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDthYztl6TrnoDroZwgMTEx!5e0!3m2!1sko!2skr!4v1661153125692!5m2!1sko!2skr",
		},
		Type:        storeType,
		Title:       "메이커2",
		Description: "메이커2 is ...",
	})
	// ..
}

func Init() {
	initHighPublic()
	// initShirtRoom()
	// initHobba()
	// ..
}
