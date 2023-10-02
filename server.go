package BWG

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	httpServer *fiber.App
}

func (s *Server) Run(port string, handler *fiber.App) error {
	s.httpServer = handler
	return s.httpServer.Listen(":" + port)
}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.ShutdownWithContext(ctx)
}
