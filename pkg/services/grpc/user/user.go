package user

import (
	"context"

	"github.com/SSH-Management/protobuf/server/users"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct{}

func (u *UserService) Create(ctx context.Context, dto *users.CreateUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
