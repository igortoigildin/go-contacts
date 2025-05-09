package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/igortoigildin/go-contacts/auth/pkg/proto"
)

type Config struct {
	GRPCPort        string
	GRPCGatewayPort string

	ChainUnaryInterceptors []grpc.UnaryServerInterceptor
	UnaryInterceptors      []grpc.UnaryServerInterceptor
}

type Controllers struct {
	pb.AuthServiceServer
}

type Server struct {
	Controllers

	grpc struct {
		lis    net.Listener
		server *grpc.Server
	}

	grpcGateway struct {
		lis    net.Listener
		server *http.Server
	}
}

// New - returns *Server
func New(ctx context.Context, cfg Config, svcs Controllers) (*Server, error) {
	srv := &Server{
		Controllers: svcs,
	}

	// grpc
	{
		// middlewares
		grpcServerOptions := unaryInterceptorsToGrpcServerOptions(cfg.UnaryInterceptors...)
		grpcServerOptions = append(grpcServerOptions,
			grpc.ChainUnaryInterceptor(cfg.ChainUnaryInterceptors...),
		)

		grpcServer := grpc.NewServer(grpcServerOptions...)
		// router
		//pb.RegisterAuthServiceServer(grpcServer, srv)

		pb.RegisterAuthServiceServer(grpcServer, srv)

		reflection.Register(grpcServer)

		lis, err := net.Listen("tcp", cfg.GRPCPort)
		if err != nil {
			return nil, err
		}

		srv.grpc.lis = lis
		srv.grpc.server = grpcServer
	}

	// grpc gateway
	{
		// router
		mux := runtime.NewServeMux()
		if err := pb.RegisterAuthServiceHandlerServer(context.Background(), mux, srv); err != nil {
			return nil, fmt.Errorf("failed to register auth service handler server: %w", err)
		}

		// middlewares
		// ...

		httpServer := &http.Server{Handler: mux}

		lis, err := net.Listen("tcp", cfg.GRPCGatewayPort)
		if err != nil {
			return nil, fmt.Errorf("server: failed to listen: %v", err)
		}

		srv.grpcGateway.lis = lis
		srv.grpcGateway.server = httpServer
	}

	return srv, nil
}

// Run - serve grpc and grpc gateway
func (s *Server) Run(ctx context.Context) error {
	group := errgroup.Group{}

	group.Go(func() error {
		log.Println("start serve grpc", s.grpc.lis.Addr())
		if err := s.grpc.server.Serve(s.grpc.lis); err != nil {
			return fmt.Errorf("server: serve grpc: %w", err)
		}
		return nil
	})

	group.Go(func() error {
		log.Println("start serve grpc gateway", s.grpcGateway.lis.Addr())
		return fmt.Errorf("server: serve grpc gateway: %w", s.grpcGateway.server.Serve(s.grpcGateway.lis))
	})

	return group.Wait()
}

func unaryInterceptorsToGrpcServerOptions(interceptors ...grpc.UnaryServerInterceptor) []grpc.ServerOption {
	opts := make([]grpc.ServerOption, 0, len(interceptors))
	for _, interceptor := range interceptors {
		opts = append(opts, grpc.UnaryInterceptor(interceptor))
	}
	return opts
}
