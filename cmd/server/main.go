package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Archisman-Mridha/chat-service/api"
	"github.com/Archisman-Mridha/chat-service/api/proto/generated"
	"github.com/Archisman-Mridha/chat-service/config"
	"github.com/Archisman-Mridha/chat-service/constants"
	"github.com/Archisman-Mridha/chat-service/internal/usecases"
	"github.com/Archisman-Mridha/chat-service/pkg/grpc"
	"github.com/Archisman-Mridha/chat-service/pkg/healthcheck"
	"github.com/Archisman-Mridha/chat-service/pkg/logger"
	"github.com/Archisman-Mridha/chat-service/pkg/postgres"
	"github.com/Archisman-Mridha/chat-service/pkg/redis"
	"github.com/Archisman-Mridha/chat-service/pkg/utils"
	"golang.org/x/sync/errgroup"
)

var configFilePath string

// Parses command line flags.
func init() {
	flagSet := flag.NewFlagSet("", flag.ExitOnError)

	flagSet.StringVar(&configFilePath, "config-file", "/config.yaml", "Path to the config file")

	envPrefix := "CHAT_MICROSERVICE_"
	flagSet.VisitAll(utils.CreateGetFlagOrEnvValueFn(envPrefix))

	cmdArgs := os.Args[1:]
	if err := flagSet.Parse(cmdArgs); err != nil {
		log.Fatalf("Failed parsing command line flags : %v", err)
	}
}

func main() {
	// When the program receives any interruption / SIGKILL / SIGTERM signal, the cancel function is
	// automatically invoked. The cancel function is responsible for freeing all the resources
	// associated with the context.
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
	)
	defer cancel()

	if err := run(ctx); err != nil {
		slog.Error(err.Error())

		cancel()
		time.Sleep(constants.RESOURCE_CLEANUP_TIMEOUT) // Give some time for remaining resources to be freed.

		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// Parse config file.
	config, err := utils.ParseConfigFile[config.Config](configFilePath)
	if err != nil {
		return err
	}

	waitGroup, waitGroupCtx := errgroup.WithContext(ctx)

	// Initialize logger, tracer and HTTP metrics server.
	logger.InitLogger()
	traceExporter := utils.InitTracer(ctx, constants.SERVICE_NAME, constants.SERVICE_VERSION, config.JaegerURL)
	metricsServer := utils.InitMetricsServer(waitGroup, waitGroupCtx, config.MetricsServerPort)

	var (
		postgresAdapter = postgres.NewDatabaseAdapter(config.PostgresURL)
		redisAdapter    = redis.NewKVStoreAdapter(config.RedisURL)

		healthcheckables = []healthcheck.Healthcheckable{
			postgresAdapter,
			redisAdapter,
		}

		usecases = usecases.NewUsecases(postgresAdapter, redisAdapter)
	)

	// Run gRPC server.
	server := grpc.CreateGRPCServer(healthcheckables)
	generated.RegisterChatServiceServer(server, api.NewChatServiceGRPCAPI(usecases))
	waitGroup.Go(func() error {
		return grpc.RunGRPCServer(server, config.GRPCServerPort)
	})

	// Handle shutdown gracefully.
	waitGroup.Go(func() error {
		/*
			The returned channel gets closed when either of this happens :

				(1) A program termination signal is received, because of which the parent context's done
					channel gets closed.

				(2) Any of the go-routines registered under this wait-group, finishes running.
		*/
		<-waitGroupCtx.Done()
		slog.Info("Gracefully shutting down program. Cleaning up resources")

		// Stop the gRPC server from accepting new connections and RPCs and block until all the pending
		// RPCs are finished.
		server.GracefulStop()
		slog.Info("Shutdown gRPC server")

		if err := metricsServer.Close(); err != nil {
			slog.Error("Failed shutting down HTTP metrics server", logger.Error(err))
		}
		slog.Info("Shutdown HTTP metrics server")

		if err := traceExporter.Shutdown(context.Background()); err != nil {
			slog.Error("Failed shutting down trace exporter", logger.Error(err))
		}
		slog.Info("Shutdown trace exporter")

		return nil
	})

	return waitGroup.Wait()
}
