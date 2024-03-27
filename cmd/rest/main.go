package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	provider, err := Initialize(ctx)
	if err != nil {
		fmt.Println("start api error:", err)
	}

	provider.ListenAndServe()
}
