package config

import (
	"log"
	"math"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	UsersGrpc() *grpc
	ItemsGrpc() *grpc
}

func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv failed: %v", err)
	}

	return &config{
		app: &app{
			url:     envMap["APP_URL"],
			name:    envMap["APP_NAME"],
			version: envMap["APP_VERSION"],
			readTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_READ_TIMEOUT"])
				if err != nil {
					log.Fatalf("load read timeout failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			writeTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_WRITE_TIMEOUT"])
				if err != nil {
					log.Fatalf("load read timeout failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			bodyLimit: func() int {
				b, err := strconv.Atoi(envMap["APP_BODY_LIMIT"])
				if err != nil {
					log.Fatalf("load body limit failed: %v", err)
				}
				return b
			}(),
		},
		db: &db{
			url:      envMap["DB_URL"],
			database: envMap["DB_DATABASE"],
		},
		usersGrpc: &grpc{
			url: envMap["USERS_GRPC_URL"],
		},
		itemsGrpc: &grpc{
			url: envMap["ITEMS_GRPC_URL"],
		},
	}
}

func (cfg *config) App() IAppConfig {
	return cfg.app
}
func (cfg *config) Db() IDbConfig {
	return cfg.db
}
func (cfg *config) UsersGrpc() *grpc {
	return cfg.usersGrpc
}
func (cfg *config) ItemsGrpc() *grpc {
	return cfg.itemsGrpc
}

type config struct {
	app       *app
	db        *db
	usersGrpc *grpc
	itemsGrpc *grpc
}

type IAppConfig interface {
	Url() string
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
}

type app struct {
	url          string
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int
}

func (a *app) Url() string                 { return a.url }
func (a *app) Name() string                { return a.name }
func (a *app) Version() string             { return a.version }
func (a *app) ReadTimeout() time.Duration  { return a.readTimeout }
func (a *app) WriteTimeout() time.Duration { return a.writeTimeout }
func (a *app) BodyLimit() int              { return a.bodyLimit }

type IDbConfig interface {
	Url() string
	Database() string
}

type db struct {
	url      string
	database string
}

func (d *db) Url() string      { return d.url }
func (d *db) Database() string { return d.database }

type grpc struct {
	url string
}

func (g *grpc) Url() string { return g.url }
