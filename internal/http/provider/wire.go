//go:generate wire
//go:build wireinject
// +build wireinject

package provider

import (
	"context"

	"capital-gains-api/internal/domain/service/operation"
	"capital-gains-api/internal/domain/service/tax"
	"capital-gains-api/internal/http/controller"
	"capital-gains-api/internal/http/handler"
	"capital-gains-api/internal/http/router"

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
