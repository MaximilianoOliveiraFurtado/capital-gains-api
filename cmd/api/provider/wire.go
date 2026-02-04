//go:generate wire
//go:build wireinject
// +build wireinject

package provider

import (
	"context"

	"capital-gains-api/cmd/rest/controller"
	"capital-gains-api/cmd/rest/handler"
	"capital-gains-api/cmd/rest/router"
	"capital-gains-api/internal/entity"
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

var entities = wire.NewSet(
	entity.NewFinstate,
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
		entities,
		services,
		controllers,
		cfg,
	)
	return Wire{}, nil
}
