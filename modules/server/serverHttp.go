package server

import (
	"encoding/json"

	"github.com/Rayato159/mytems-microservices/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type IServerHttp interface {
	RunUsersHttp()
	RunItemsHttp()
	HttpServer() *httpServer
}

type httpServer struct {
	app *fiber.App
	cfg config.IConfig
	db  *mongo.Client
}

func NewHttpServer(cfg config.IConfig, db *mongo.Client) IServerHttp {
	return &httpServer{
		cfg: cfg,
		db:  db,
		app: fiber.New(fiber.Config{
			AppName:      cfg.App().Name(),
			BodyLimit:    cfg.App().BodyLimit(),
			ReadTimeout:  cfg.App().ReadTimeout(),
			WriteTimeout: cfg.App().WriteTimeout(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
	}
}

func (s *httpServer) HttpServer() *httpServer { return s }
