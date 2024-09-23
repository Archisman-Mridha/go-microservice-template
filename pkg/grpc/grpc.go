package grpc

import (
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/Archisman-Mridha/chat-service/constants"
	"github.com/Archisman-Mridha/chat-service/pkg/healthcheck"
	"github.com/Archisman-Mridha/chat-service/pkg/logger"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"

	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	/*
	  The WASI preview 1 specification has partial support for socket networking, preventing a large
	  class of Go applications from running when compiled to WebAssembly with GOOS=wasip1. Extensions
	  to the base specifications have been implemented by runtimes to enable a wider range of
	  programs to be run as WebAssembly modules.

	  Where possible, the package offers the ability to automatically configure the network stack via
	  init functions called on package imports.

	  When imported, this package alter the default configuration to install a dialer function
	  implemented on top of the WASI socket extensions. When compiled to other targets, the import
	  of those packages does nothing.

	  REFER : https://github.com/dev-wasm/dev-wasm-go.
	*/
	_ "github.com/stealthrocket/net/http"
)

// Creates a gRPC server. Sets up healthcheck mechanism and server reflection. The server instance
// is then returned.
func CreateGRPCServer(healthcheckables []healthcheck.Healthcheckable) *grpc.Server {
	server := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	reflection.Register(server)

	healthcheckServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthcheckServer)
	go func() {
		var healthStatus grpc_health_v1.HealthCheckResponse_ServingStatus

		for {
			if err := healthcheck.Healthcheck(healthcheckables); err != nil {
				healthStatus = grpc_health_v1.HealthCheckResponse_NOT_SERVING
				slog.Error("Healthcheck failed", logger.Error(err))
			} else {
				healthStatus = grpc_health_v1.HealthCheckResponse_SERVING
			}
			healthcheckServer.SetServingStatus("", healthStatus)

			time.Sleep(constants.HEALTHCHECK_FREQUENCY)
		}
	}()

	return server
}

// Creates a TCP listener at the given address and uses it to run the given gRPC server instance.
func RunGRPCServer(server *grpc.Server, port int) error {
	address := fmt.Sprintf("0.0.0.0:%d", port)
	tcpListener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed creating TCP listener : %v", err)
	}

	slog.Info("Starting gRPC server", slog.String("address", address))
	if err := server.Serve(tcpListener); err != nil {
		return fmt.Errorf("gRPC server error occurred : %v", err)
	}
	return nil
}
