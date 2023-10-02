package handler

import (
	"BWG/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoute() *fiber.App {
	router := fiber.New()

	api := router.Group("/api")
	api.Get("/test", h.test)
	//api.Post("/withdraw", h.PostWithdrawHandler)
	//api.Post("/transfer", h.PostTransferHandler)
	//api.Get("/balance", h.GetBalanceHandler)

	return router
}

func (h Handler) test(c *fiber.Ctx) error {
	return c.SendString("I'm a GET request!")
}
