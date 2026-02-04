//go:generate wire
//go:build wireinject
// +build wireinject

package provider

import (
	"context"

	"capital-gains-api/cmd/api/controller"
	"capital-gains-api/cmd/api/handler"
	"capital-gains-api/cmd/api/router"
	"capital-gains-api/internal/service/operation"
	"capital-gains-api/internal/service/tax"

	"github.com/google/wire"
)

type Wire struct {
	Router *router.Router
}

func newWire(httpRouter *router.Router) Wire {
	return Wire{
		Router: httpRouter,
	}
}

var cfg = wire.NewSet(
	newWire,
	router.NewHTTPRouter,
)

var handlers = wire.NewSet(
	handler.NewHandler,
)

var services = wire.NewSet(
	tax.NewService,
	operation.NewService,
)

var controllers = wire.NewSet(
	controller.NewOperationTaxController,
)

func Initialize(ctx context.Context) (Wire, error) {
	wire.Build(
		handlers,
		services,
		controllers,
		cfg,
	)
	return Wire{}, nil
}
