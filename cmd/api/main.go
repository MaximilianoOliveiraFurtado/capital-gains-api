package main

import (
	"context"
	"fmt"

	"capital-gains-api/internal/http/provider"
)

func main() {
	ctx := context.Background()

	app, err := provider.Initialize(ctx)
	if err != nil {
		fmt.Println("start app error:", err)
	}

	err = app.Router.ListenAndServe()
	if err != nil {
		fmt.Println("start server error:", err)
	}

}
