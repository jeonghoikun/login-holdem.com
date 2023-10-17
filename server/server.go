package server

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/template/html/v2"
	"github.com/jeonghoikun/login-holdem.com/site"
)

type port uint32

func (p *port) String() string { return fmt.Sprintf(":%d", *p) }

type Server struct {
	port *port
	app  *fiber.App
}

type engineFunc struct{}

func (*engineFunc) time() time.Time { return time.Now() }

func (*engineFunc) withHost(s string) string {
	return fmt.Sprintf("https://%s%s", site.Config.Domain, s)
}

func (*engineFunc) randFileNumber() int {
	min := 0
	max := 1000
	return rand.Intn(max-min) + min
}

func engine() *html.Engine {
	e := html.New("./views", ".html")
	e.Reload(true)
	ef := &engineFunc{}
	e.AddFunc("Time", ef.time)
	e.AddFunc("WithHost", ef.withHost)
	e.AddFunc("RandFileNumber", ef.randFileNumber)
	return e
}

func New(portNumber uint32) *Server {
	p := port(portNumber)
	app := fiber.New(fiber.Config{
		AppName:      site.Config.Domain,
		ServerHeader: site.Config.Domain,
		Views:        engine(),
	})
	return &Server{port: &p, app: app}
}

func (s *Server) set() {
	s.app.Static("/static", "./static")
}

func (s *Server) middlewares() {
	s.app.Use("/",
		compress.New(compress.Config{Level: compress.Level(2)}),
		bindSiteConfig,
	)
}

func (s *Server) routes() {
	handleIndex(s.app.Group("/"))
}

func (s *Server) Run() error {
	s.set()
	s.middlewares()
	s.routes()
	return s.app.Listen(s.port.String())
}

type Author struct {
	Name        string
	ProfilePath string
}

type PageConfig struct {
	Path          string
	Author        *Author
	Title         string
	Description   string
	Keywords      string
	PhoneNumber   string
	DatePublished time.Time
	DateModified  time.Time
	ThumbnailPath string
}
