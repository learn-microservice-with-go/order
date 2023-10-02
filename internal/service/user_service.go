package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	proto "github.com/learn-microservice-with-go/user_microservice/api/service"
	model "github.com/learn-microservice-with-go/user_microservice/internal/model"
	"github.com/learn-microservice-with-go/user_microservice/internal/utils"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	Db    *gorm.DB
	Mongo *mongo.Client
	Rds   *redis.Client
	proto.UnimplementedUserServiceServer
}

// BatchGetUsers implements proto.UserServiceServer.
func (s *Server) BatchGetUsers(context.Context, *proto.BatchGetUsersRequest) (*proto.BatchGetUsersReply, error) {
	return &proto.BatchGetUsersReply{}, nil
}

// CreateUser implements proto.UserServiceServer.
func (s *Server) CreateUser(context.Context, *proto.CreateUserRequest) (*proto.CreateUserReply, error) {
	return &proto.CreateUserReply{}, nil
}

// Login implements proto.UserServiceServer.
func (s *Server) Login(context.Context, *proto.LoginRequest) (*proto.LoginReply, error) {
	return &proto.LoginReply{}, nil
}

// Logout implements proto.UserServiceServer.
func (s *Server) Logout(context.Context, *proto.LogoutRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// Register implements proto.UserServiceServer.
func (s *Server) Register(context.Context, *proto.RegisterRequest) (*proto.RegisterReply, error) {
	return &proto.RegisterReply{}, nil
}

// UpdatePassword implements proto.UserServiceServer.
func (s *Server) UpdatePassword(context.Context, *proto.UpdatePasswordRequest) (*proto.UpdatePasswordReply, error) {
	return &proto.UpdatePasswordReply{}, nil
}

// UpdateUser implements proto.UserServiceServer.
func (s *Server) UpdateUser(context.Context, *proto.UpdateUserRequest) (*proto.UpdateUserReply, error) {
	return &proto.UpdateUserReply{}, nil
}

func (s *Server) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserReply, error) {
	log.Printf("Receive message body from client]: %s", strconv.FormatInt(in.Id, 10))

	// Redis
	_, err := s.Rds.Set(ctx, uuid.NewString(), utils.GenerateVerificationCode(6), 10*time.Minute).Result()
	if err != nil {
		panic(err)
	}

	// MySQL
	var offices []model.Office
	if err := s.Db.Raw("SELECT * FROM offices").Scan(&offices).Error; err != nil {
		log.Fatalf("Failed to query database: %v", err)
	}
	mysqlData, err := json.Marshal(offices)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(mysqlData))

	// Mongo
	collection := s.Mongo.Database("blog").Collection("announcement")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	mongoData, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(mongoData))

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
