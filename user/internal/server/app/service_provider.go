package app

import (
	"context"
	"log"

	userApi "github.com/igortoigildin/go-contacts/user/internal/server/api"
	"github.com/igortoigildin/go-contacts/user/internal/server/config"
	"github.com/igortoigildin/go-contacts/user/internal/server/service"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig

	userService *service.UserService

	userImpl *userApi.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("error creating gRPC config: %v", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) UserService(ctx context.Context) *service.UserService {
	if s.userService == nil {
		s.userService = service.New(ctx)
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *userApi.Implementation {
	if s.userImpl == nil {
		s.userImpl = userApi.NewImplementation(s.UserService(context.Background()))
	}

	return s.userImpl
}
