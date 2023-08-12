package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type storeHandler struct{}

// GET /store/:do/:si/:dong/:storeType/:title
func (*storeHandler) page(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("store page")
}

// BaseURL = /store
func handleStore(r fiber.Router) {
	h := &storeHandler{}
	r.Get("/:do/:si/:dong/:storeType/:title", h.page)
}
