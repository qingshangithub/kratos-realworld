package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/errors"
)

func (s *RealWorldService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	if len(req.User.Email) == 0 {
		return nil, errors.NewHTTPError(422, "email", "cannot be empty")
	}

	return &pb.UserReply{
		User: &pb.UserReply_User{
			Username: "boom",
		},
	}, nil
}
func (s *RealWorldService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &pb.UserReply{
		User: &pb.UserReply_User{
			Username: u.Username,
			Token:    u.Token,
		},
	}, nil
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *emptypb.Empty) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
