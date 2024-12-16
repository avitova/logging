package main

import (
	"context"
	"log/slog"

	"github.com/lzap/strc"
)

func subProcess(ctx context.Context) {
	span, ctx := strc.StartContext(ctx, "subProcess")
	defer span.End()
}

func process(ctx context.Context) {
	span, ctx := strc.StartContext(ctx, "process")
	defer span.End()

	subProcess(ctx)
}

func main() {
	// tracing logs via DebugLevel by default
	strc.Level = slog.LevelInfo

	span, ctx := strc.StartContext(context.Background(), "main")
	defer span.End()

	process(ctx)
}
