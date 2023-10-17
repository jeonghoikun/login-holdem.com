package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonghoikun/login-holdem.com/site"
)

type indexHandler struct{}

// GET /
func (*indexHandler) index(c *fiber.Ctx) error {
	m := fiber.Map{
		"Page": &PageConfig{
			Path: c.Path(),
			Author: &Author{
				Name:        site.Config.Author,
				ProfilePath: "/static/img/site/author/profile.png",
			},
			Title:         site.Config.Title,
			Description:   site.Config.Description,
			Keywords:      site.Config.Keywords.String(),
			PhoneNumber:   site.Config.PhoneNumber,
			DatePublished: site.Config.DatePublished,
			DateModified:  site.Config.DateModified,
			ThumbnailPath: "/static/img/site/thumbnail/thumb.png",
		},
	}
	return c.Status(http.StatusOK).Render("index", m, "layout/index")
}

// GET /robots.txt
func (*indexHandler) robots(c *fiber.Ctx) error {
	var ss []string
	ss = append(ss, "User-agent: *")
	ss = append(ss, "Allow: /")
	ss = append(ss, fmt.Sprintf("Sitemap: %s/sitemap.xml", site.Config.Domain))
	return c.Status(http.StatusOK).SendString(strings.Join(ss, "\n"))
}

// GET /sitemap.xml
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

	ss = append(ss, `</urlset>`)
	c.Response().Header.Add("Content-Type", "application/xml")
	return c.Status(http.StatusOK).SendString(strings.Join(ss, "\n"))
}

// BaseURL = /
func handleIndex(r fiber.Router) {
	h := &indexHandler{}
	r.Get("/", h.index)
	r.Get("/robots.txt", h.robots)
	r.Get("/sitemap.xml", h.sitemap)
}
