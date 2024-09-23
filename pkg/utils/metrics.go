package utils

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
)

func InitMetricsServer(waitGroup *errgroup.Group, waitGroupCtx context.Context, port int) *http.Server {
	server := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", port),
	}

	http.Handle("/metrics", promhttp.Handler())

	waitGroup.Go(func() error {
		slog.Info("Starting metrics HTTP server", slog.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	return server
}
