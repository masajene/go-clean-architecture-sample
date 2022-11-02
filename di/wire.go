//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"go_api_boilerplate/adapter/handler"
)

func InitUserApp() *handler.UserHandler {
	wire.Build(handler.WireSet)
	return nil
}
