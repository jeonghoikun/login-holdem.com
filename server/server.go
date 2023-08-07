package server

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/template/html/v2"
)

type port uint32

func (p *port) String() string { return fmt.Sprintf(":%d", *p) }

type Server struct {
	port *port
	app  *fiber.App
}

type engineFunc struct{}

func (*engineFunc) time() time.Time { return time.Now() }

func engine() *html.Engine {
	e := html.New("./views", ".html")
	e.Reload(true)
	ef := &engineFunc{}
	e.AddFunc("Time", ef.time)
	return e
}

func New(portNumber uint32) *Server {
	p := port(portNumber)
	app := fiber.New(fiber.Config{
		AppName:      "",
		ServerHeader: "",
		Views:        engine(),
	})
	return &Server{port: &p, app: app}
}

func (s *Server) set() {
	s.app.Static("/static", "./static")
}

func (s *Server) middlewares() {
	s.app.Use("/", compress.New(compress.Config{Level: compress.Level(2)}))
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
