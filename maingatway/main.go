package main

//
//import (
//	"context"
//	"flag"
//	"net/http"
//	"user_growth/pb"
//
//	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"google.golang.org/grpc/grpclog"
//)
//
//var (
//	// command-line options:
//	// gRPC server endpoint
//	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8080", "gRPC server endpoint")
//)
//
//func run() error {
//	ctx := context.Background()
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel()
//
//	// Register gRPC server endpoint
//	// Note: Make sure the gRPC server is running properly and accessible
//	mux := runtime.NewServeMux()
//	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
//	err := pb.RegisterUserCoinHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
//	if err != nil {
//		return err
//	}
//	err = pb.RegisterUserGradeHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
//	if err != nil {
//		return err
//	}
//
//	// Start HTTP server (and proxy calls to gRPC server endpoint)
//	return http.ListenAndServe(":8081", mux)
//}
//
//func main() {
//	flag.Parse()
//
//	if err := run(); err != nil {
//		grpclog.Fatal(err)
//	}
//}

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"user_growth/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8080", "gRPC server endpoint")
	httpPort           = flag.String("http-port", "8081", "HTTP server port")
)

func main() {
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if err := run(logger); err != nil {
		logger.Fatal("Failed to run server", zap.Error(err))
	}
}

func run(logger *zap.Logger) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := pb.RegisterUserCoinHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return fmt.Errorf("failed to register UserCoin handler: %v", err)
	}
	if err := pb.RegisterUserGradeHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return fmt.Errorf("failed to register UserGrade handler: %v", err)
	}

	gwServer := &http.Server{
		Addr:    ":" + *httpPort,
		Handler: cors(tracing(logging(mux, logger))),
	}

	logger.Info("Starting gateway server", zap.String("port", *httpPort))
	go func() {
		if err := gwServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to listen and serve", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutting down gateway server...")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := gwServer.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Gateway server exited")
	return nil
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func logging(h http.Handler, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		logger.Info("Request processed",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Duration("duration", time.Since(start)),
		)
	})
}

func tracing(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add distributed tracing logic here, e.g., using OpenTelemetry
		h.ServeHTTP(w, r)
	})
}
