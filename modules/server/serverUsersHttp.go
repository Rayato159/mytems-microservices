package server

import (
	"log"
	"os"
	"os/signal"
)

func (s *httpServer) RunUsersHttp() {
	// Middlewares

	// Modules

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	// Listen to host:port
	log.Printf("%s is starting on %s", s.cfg.App().Name(), s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())
}
