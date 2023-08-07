package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type indexHandler struct{}

func (*indexHandler) index(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("index")
}

func handleIndex(r fiber.Router) {
	h := &indexHandler{}
	r.Get("/", h.index)
}
