package controllers

import (
	"context"

	pb "github.com/igortoigildin/go-contacts/auth/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controller) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// 1. validate request
	err := c.Validator.Validate(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	// 2. convert delivery models to domain models/DTO
	registerDTO := newRegisterDTOFromRegisterRequest(req)

	// 3. call usecase
	username, err := c.AuthUsecase.Register(ctx, registerDTO)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register: %v", err)
	}

	// 4. convert domain models to delivery models
	response := &pb.RegisterResponse{
		Username: username,
	}

	return response, nil
}

func (c *Controller) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// 1. validate request
	err := c.Validator.Validate(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	// 2. convert delivery models to domain models/DTO
	loginDTO := newLoginDTOFromLoginRequest(req)

	// 3. call usecase
	token, err := c.AuthUsecase.Login(ctx, loginDTO)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to login: %v", err)
	}

	// 4. convert domain models to delivery models
	response := &pb.LoginResponse{
		Token: token,
	}

	return response, nil
}

func (c *Controller) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.SuccessResponse, error) {
	// 1. validate request
	err := c.Validator.Validate(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	// 2. convert delivery models to domain models/DTO
	logoutDTO := newLogoutDTOFromLogoutRequest(req)

	// 3. call usecase
	_, err = c.AuthUsecase.Logout(ctx, logoutDTO)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to logout: %v", err)
	}

	// 4. convert domain models to delivery models
	response := &pb.SuccessResponse{
		Message: "Logout successful",
	}

	return response, nil
}

func (c *Controller) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.SuccessResponse, error) {
	// 1. validate request
	err := c.Validator.Validate(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	// 2. convert delivery models to domain models/DTO
	verifyDTO := newVerifyDTOFromVerifyRequest(req)

	// 3. call usecase
	_, err = c.AuthUsecase.Verify(ctx, verifyDTO)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify: %v", err)
	}

	// 4. convert domain models to delivery models
	response := &pb.SuccessResponse{
		Message: "Verification successful",
	}

	return response, nil
}
