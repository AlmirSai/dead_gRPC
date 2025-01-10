// internal/grpc/auth/server.go
package authgrpc

import (
	"context"
	"errors"

	"grpc-service-ref/internal/services/auth"
	"grpc-service-ref/internal/storage"

	ssov1 "github.com/JustSkiv/protos/gen/go/sso"
	"github.com/docker/distribution/registry/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServiceServer
	auth Auth
}

type Auth interface {
    Login(
        ctx context.Context,
        email string,
        password string,
        appID int,
    ) (token string, err error)
    RegisterNewUser(
        ctx context.Context,
        email string,
        password string,
    ) (userID int64, err error)
}

func Register(gRPCServer *grpc.Server, auth Auth) {  
    ssov1.RegisterAuthServer(gRPCServer, &serverAPI{auth: auth})  
}
func (s *serverAPI) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "Email is required")
	}
	
	if in.password == "" {
		return nil, status.Error(codes.InvalidArgument, "Password is required")
	}

	if in.GetAppId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "AppID is required")
	}

	token, err := s.auth.Login(ctx, in.GetEmail(), in.GetPassword(), int(in.GetAppId()))
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredential) {
			return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
		}

		return nil, status.Error(codes.Internal, "Internal error. Please try again later")
	}

	return &ssov1.LoginResponse{Token: token}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.RegisterResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "Email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "Password is required")
	}

	uid, err := s.auth.RegisterNewUser(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "User already exists")
		}

		return nil, status.Error(codes.Internal, "Internal error. Please try again later")
	}

	return &ssov1.RegisterResponse{UserID: uid}, nil
}
