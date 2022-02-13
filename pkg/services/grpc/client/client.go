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

var _ clients.ClientServiceServer = &Service{}

type Service struct {
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
) *Service {
	return &Service{
		serverPublicKey:  serverPublicKey,
		logger:           logger,
		serverRepository: serverRepository,
		userRepository:   userRepository,
		validator:        validator,
	}
}

func (cl *Service) Create(ctx context.Context, req *clients.CreateClientRequest) (*clients.CreateClientResponse, error) {
	cl.logger.Debug().
		Str("name", req.Name).
		Str("ip", req.Ip).
		Str("group", req.Group).
		Msg("Client Request")

	if err := conform.Strings(&req); err != nil {
		cl.logger.Error().Err(err).
			Str("name", req.Name).
			Str("ip", req.Ip).
			Str("group", req.Group).
			Msg("error while conforming strings")
		return nil, err
	}

	if err := cl.validator.Struct(req); err != nil {
		cl.logger.Error().Err(err).
			Str("name", req.Name).
			Str("ip", req.Ip).
			Str("group", req.Group).
			Msg("validation errors")
		return nil, err
	}

	s, err := cl.serverRepository.FindByPrivateIP(ctx, req.Ip)
	if err != nil {
		if err == db.ErrNotFound {
			s, err = cl.serverRepository.Create(ctx, req)

			if err != nil {
				cl.logger.Error().Err(err).
					Str("name", req.Name).
					Str("ip", req.Ip).
					Str("group", req.Group).
					Msg("Error while create server in database")

				return nil, status.Error(codes.Internal, "Error while creating server")
			}

		} else {
			cl.logger.Error().Err(err).
				Str("name", req.Name).
				Str("ip", req.Ip).
				Str("group", req.Group).
				Msg("error while fetching server from database")

			return nil, status.Error(codes.Internal, "Error while creating server")
		}
	}

	users, err := cl.userRepository.FindByGroup(ctx, s.GroupID)
	if err != nil {
		cl.logger.Error().Err(err).
			Str("name", req.Name).
			Str("ip", req.Ip).
			Str("group", req.Group).
			Msg("error while finding server group")

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

func (cl *Service) Delete(ctx context.Context, req *clients.DeleteClientRequest) (*emptypb.Empty, error) {
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
