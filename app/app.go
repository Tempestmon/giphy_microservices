package app

import (
	apiservice "giphy_microservices/api_service"

	"github.com/google/wire"
)

type App struct {
	HttpServer *apiservice.Server
}

func InitApp() *App {
	wire.Build(apiservice.NewServer, wire.Struct(new(App), "*"))
	return &App{}
}
