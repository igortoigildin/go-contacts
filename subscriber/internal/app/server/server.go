package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/igortoigildin/go-contacts/subscriber/pkg/proto"
)

type Config struct {
	GRPCPort string

	ChainUnaryInterceptors []grpc.UnaryServerInterceptor
	UnaryInterceptors      []grpc.UnaryServerInterceptor
}

type Controllers struct {
	pb.SubscriberServiceServer
}

type Server struct {
	Controllers

	grpc struct {
		lis    net.Listener
		server *grpc.Server
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

		pb.RegisterSubscriberServiceServer(grpcServer, srv)

		reflection.Register(grpcServer)

		lis, err := net.Listen("tcp", cfg.GRPCPort)
		if err != nil {
			return nil, err
		}

		srv.grpc.lis = lis
		srv.grpc.server = grpcServer
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

	return group.Wait()
}

func unaryInterceptorsToGrpcServerOptions(interceptors ...grpc.UnaryServerInterceptor) []grpc.ServerOption {
	opts := make([]grpc.ServerOption, 0, len(interceptors))
	for _, interceptor := range interceptors {
		opts = append(opts, grpc.UnaryInterceptor(interceptor))
	}
	return opts
}

// Subscriber routes
// e.POST("/api/v1/subscriber/request", subscriber.MakeFriendRequestHandler)
// e.POST("/api/v1/subscriber/accept", subscriber.AcceptFriendHandler)
// e.POST("/api/v1/subscriber/reject", subscriber.RejectFriendHandler)
// e.DELETE("/api/v1/subscriber", subscriber.RemoveFriendHandler)
// e.GET("/api/v1/subscriber/:id", subscriber.GetFriendsHandler)

// e.GET("/health/ready", func(c echo.Context) error {
// 	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
// })
