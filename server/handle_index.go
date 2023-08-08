package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonghoikun/webserver/site"
	"github.com/jeonghoikun/webserver/store"
)

type indexHandler struct{}

func (*indexHandler) index(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("index")
}

func (*indexHandler) robots(c *fiber.Ctx) error {
	var ss []string
	ss = append(ss, "User-agent: *")
	ss = append(ss, "Allow: /")
	ss = append(ss, fmt.Sprintf("Sitemap: %s/sitemap.xml", site.Config.Domain))
	return c.Status(http.StatusOK).SendString(strings.Join(ss, "\n"))
}

func (*indexHandler) sitemap(c *fiber.Ctx) error {
	var ss []string
	ss = append(ss, `<?xml version="1.0" encoding="UTF-8"?>`)
	ss = append(ss, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)

	host := fmt.Sprintf("https://%s", site.Config.Domain)
	dateModified := site.Config.DateModified.Format(time.RFC3339)

	// index
	ss = append(ss, `<url>`)
	ss = append(ss, fmt.Sprintf(`<loc>%s/</loc>`, host))
	ss = append(ss, fmt.Sprintf(`<lastmod>%s</lastmod>`, dateModified))
	ss = append(ss, `</url>`)

	// categories
	// TODO:

	// stores
	for _, s := range store.ListAllStores() {
		ss = append(ss, `<url>`)
		do := url.QueryEscape(s.Location.Do)
		si := url.QueryEscape(s.Location.Si)
		dong := url.QueryEscape(s.Location.Dong)
		storeType := url.QueryEscape(s.Type)
		storeTitle := url.QueryEscape(s.Title)
		ss = append(ss, fmt.Sprintf(`<loc>%s/%s/%s/%s/%s/%s</loc>`, host, do, si, dong, storeType, storeTitle))
		dateModified = s.DateModified.Format(time.RFC3339)
		ss = append(ss, fmt.Sprintf(`<lastmod>%s</lastmod>`, dateModified))
		ss = append(ss, `</url>`)
	}

	ss = append(ss, `</urlset>`)
	c.Response().Header.Add("Content-Type", "application/xml")
	return c.Status(http.StatusOK).SendString(strings.Join(ss, "\n"))
}

func handleIndex(r fiber.Router) {
	h := &indexHandler{}
	r.Get("/", h.index)
	r.Get("/robots.txt", h.robots)
	r.Get("/sitemap.xml", h.sitemap)
}
