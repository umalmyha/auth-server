package grpc

import (
	"context"
	pb "github.com/umalmyha/auth-server/proto"
)

type Handler struct {
	pb.UnimplementedAuthServiceServer
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return &pb.SignUpResponse{}, nil
}
func (h *Handler) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	return &pb.SignInResponse{
		AccessToken:  "",
		RefreshToken: "",
	}, nil
}

func (h *Handler) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	return &pb.RefreshResponse{
		AccessToken:  "",
		RefreshToken: "",
	}, nil
}

func (h *Handler) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return &pb.LogoutResponse{}, nil
}
