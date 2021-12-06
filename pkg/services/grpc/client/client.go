package client

import (
	"context"

	"github.com/SSH-Management/protobuf/common"
	"github.com/SSH-Management/protobuf/server/clients"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
	"github.com/SSH-Management/server/pkg/repositories/server"
	userrepo "github.com/SSH-Management/server/pkg/repositories/user"
)

var _ clients.ClientServiceServer = &ClientService{}

type ClientService struct {
	clients.UnimplementedClientServiceServer
	logger           *log.Logger
	serverRepository server.Interface
	userRepository   userrepo.Interface
	validator        *validator.Validate

	serverPublicKey string
}

func New(
	serverPublicKey string,
	logger *log.Logger,
	serverRepository server.Interface,
	userRepository userrepo.Interface,
	validator *validator.Validate,
) *ClientService {
	return &ClientService{
		serverPublicKey:  serverPublicKey,
		logger:           logger,
		serverRepository: serverRepository,
		userRepository:   userRepository,
		validator:        validator,
	}
}

func (cl *ClientService) Create(ctx context.Context, req *clients.CreateClientRequest) (*clients.CreateClientResponse, error) {
	if err := conform.Strings(&req); err != nil {
		return nil, err
	}

	if err := cl.validator.Struct(req); err != nil {
		return nil, err
	}

	s, err := cl.serverRepository.FindByPrivateIP(ctx, req.Ip)

	if err != nil {
		if err == db.ErrNotFound {
			s, err = cl.serverRepository.Create(ctx, req)

			if err != nil {
				return nil, status.Error(codes.Internal, "Error while creating server")
			}

		} else {
			return nil, status.Error(codes.Internal, "Error while creating server")
		}
	}

	users, err := cl.userRepository.FindByGroup(ctx, s.GroupID)

	if err != nil {
		err = cl.serverRepository.Delete(ctx, s.ID)

		if err != nil {
			cl.logger.
				Error().
				Err(err).
				Msg("Error while deleting server")
		}

		return nil, status.Error(codes.Internal, "Error while creating server")
	}

	return &clients.CreateClientResponse{
		PublicKey: cl.serverPublicKey,
		Users:     mapUsers(users),
		Id:        s.ID,
	}, nil
}

func (cl *ClientService) Delete(ctx context.Context, req *clients.DeleteClientRequest) (*emptypb.Empty, error) {
	err := cl.serverRepository.Delete(ctx, req.Id)

	if err != nil {
		return nil, status.Error(codes.Internal, "Error while deleting server")
	}

	return &emptypb.Empty{}, nil
}

func mapUsers(users []models.User) []*common.LinuxUser {
	userMap := make([]*common.LinuxUser, 0, len(users))

	for _, user := range users {
		userMap = append(userMap, &common.LinuxUser{
			Username:     user.Username,
			Shell:        user.Shell,
			SystemGroups: []string{"sudo"},
			PublicKey:    user.PublicSSHKey,
		})
	}

	return userMap
}
