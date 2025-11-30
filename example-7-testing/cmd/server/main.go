package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"

	"github.com/elisardofelix/grpc-examples/example-7-testing/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := app.Run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("error running application",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	slog.Info("closing server gracefully")
}
