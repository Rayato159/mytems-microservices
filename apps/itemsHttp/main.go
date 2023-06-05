package main

import (
	"os"

	"github.com/Rayato159/mytems-microservices/config"
	"github.com/Rayato159/mytems-microservices/modules/server"
	"github.com/Rayato159/mytems-microservices/pkg/database"
)

func main() {
	path := func() string {
		if len(os.Args) == 1 {
			return ".env.items.http"
		}
		return os.Args[1]
	}()

	cfg := config.LoadConfig(path)

	client := database.DBConnect(cfg)
	defer database.DBDisconnect(client)

	db := client.Database(cfg.Db().Database())

	server.NewHttpServer(cfg, db).RunItemsHttp()
}
