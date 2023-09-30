package service

import (
	"context"
	"log"
	"strconv"

	proto "github.com/learn-frame/learn-micro-service/api/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	proto.UnimplementedUserServiceServer
}

// BatchGetUsers implements proto.UserServiceServer.
func (*Server) BatchGetUsers(context.Context, *proto.BatchGetUsersRequest) (*proto.BatchGetUsersReply, error) {
	return &proto.BatchGetUsersReply{}, nil
}

// CreateUser implements proto.UserServiceServer.
func (*Server) CreateUser(context.Context, *proto.CreateUserRequest) (*proto.CreateUserReply, error) {
	return &proto.CreateUserReply{}, nil
}

// Login implements proto.UserServiceServer.
func (*Server) Login(context.Context, *proto.LoginRequest) (*proto.LoginReply, error) {
	return &proto.LoginReply{}, nil
}

// Logout implements proto.UserServiceServer.
func (*Server) Logout(context.Context, *proto.LogoutRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// Register implements proto.UserServiceServer.
func (*Server) Register(context.Context, *proto.RegisterRequest) (*proto.RegisterReply, error) {
	return &proto.RegisterReply{}, nil
}

// UpdatePassword implements proto.UserServiceServer.
func (*Server) UpdatePassword(context.Context, *proto.UpdatePasswordRequest) (*proto.UpdatePasswordReply, error) {
	return &proto.UpdatePasswordReply{}, nil
}

// UpdateUser implements proto.UserServiceServer.
func (*Server) UpdateUser(context.Context, *proto.UpdateUserRequest) (*proto.UpdateUserReply, error) {
	return &proto.UpdateUserReply{}, nil
}

func (s *Server) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserReply, error) {
	log.Printf("Receive message body from client]: %s", strconv.FormatInt(in.Id, 10))

	return &proto.GetUserReply{
		User: &proto.User{
			Id:        1,
			Username:  "yanceyofficial",
			Email:     "yanceyofficial@gmail.com",
			Phone:     "",
			LoginAt:   1,
			Status:    proto.StatusType_Ban,
			Nickname:  "",
			Avatar:    "",
			Gender:    proto.GenderType_FEMALE,
			Birthday:  "",
			Bio:       "",
			CreatedAt: 0,
			UpdatedAt: 0,
		},
	}, nil
}
