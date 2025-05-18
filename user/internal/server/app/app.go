package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	closer "github.com/igortoigildin/go-contacts/user/internal/server/closer"
	"github.com/igortoigildin/go-contacts/user/internal/server/config"
	interceptor "github.com/igortoigildin/go-contacts/user/internal/server/interceptor"
	desc "github.com/igortoigildin/go-contacts/user/pkg/api/users"
	"golang.org/x/time/rate"

	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	cfgFileName = ".env"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing dependencies: %w", err)
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := a.runGRPCServer()
		if err != nil {
			log.Fatalf("error running gRPC server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("error running HTTP server: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return fmt.Errorf("error in %v: %w", f, err)
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.LoadFromFile(cfgFileName)
	if err != nil {
		return fmt.Errorf("error loading config from local file: %w", err)
	}

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	// 10 rpc without bufer (burst = 1) - rate limiter
	limiter := rate.NewLimiter(10, 1)

	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			interceptor.ValidateInterceptor,
			interceptor.RateLimitInterceptor(limiter),
		),
	)

	reflection.Register(a.grpcServer)

	desc.RegisterUserServiceServer(a.grpcServer, a.serviceProvider.UserImpl(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := desc.RegisterUserServiceHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts)
	if err != nil {
		return fmt.Errorf("error registering user service handler: %w", err)
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: corsMiddleware.Handler(mux),
	}

	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on: %s", a.serviceProvider.HTTPConfig().Address())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return fmt.Errorf("error serving HTTP server: %w", err)
	}

	return nil
}

func (a *App) runGRPCServer() error {
	fmt.Println("GRPC server is running on: ", a.serviceProvider.grpcConfig.Address())

	list, err := net.Listen("tcp", a.serviceProvider.grpcConfig.Address())
	if err != nil {
		return fmt.Errorf("error listening on address: %w", err)
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return fmt.Errorf("error serving gRPC server: %w", err)
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
