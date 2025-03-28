//go:build wireinject
// +build wireinject

package main

import (
	"giphy_microservices/app"

	"github.com/google/wire"
)

func InitializeApp() *app.App {
	wire.Build(app.InitApp)
	return &app.App{}
}
