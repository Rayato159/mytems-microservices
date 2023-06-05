package server

import (
	"github.com/gofiber/fiber/v2"
)

type IModule interface {
}

type module struct {
	r          fiber.Router
	httpServer *httpServer
	grpcServer *grpcServer
}

func NewModule(h *httpServer, g *grpcServer) IModule {
	return &module{
		r:          h.app.Group("v1"),
		httpServer: h,
		grpcServer: g,
	}
}
