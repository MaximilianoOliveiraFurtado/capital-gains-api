package main

import (
	"context"

	"capital-gains-api/internal/service/operation"
	"capital-gains-api/internal/service/tax"

	"github.com/google/wire"
)

type Wire struct {
	*Router
}

func newWire(httpRouter *Router) Wire {
	return Wire{
		httpRouter,
	}
}

var cfg = wire.NewSet(
	newWire,
	NewHTTPRouter,
)

var handlers = wire.NewSet(
	NewHandler,
)

var services = wire.NewSet(
	tax.NewService,
	operation.NewService,
)

var controller = wire.NewSet(
	NewOperationTaxController,
)

func Initialize(ctx context.Context) (Wire, error) {
	wire.Build(
		handlers,
		services,
		controller,
		cfg,
	)
	return Wire{}, nil
}
