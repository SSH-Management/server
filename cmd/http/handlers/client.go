package handlers

import (
	"encoding/base64"
	"net/http"
	"os"

	sdk "github.com/SSH-Management/server-sdk"
	"github.com/SSH-Management/utils"
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/repositories/server"
	userrepo "github.com/SSH-Management/server/pkg/repositories/user"
)

func CreateNewClientHandler(
	publicKeyPath string,
	logger *log.Logger,
	serverRepository server.Interface,
	userRepository userrepo.Interface,
) fiber.Handler {
	path, err := utils.GetAbsolutePath(publicKeyPath)

	if err != nil {
		logger.Fatal().
			Err(err).
			Str("path", publicKeyPath).
			Msg("Failed to get absolute path for the public key")
	}

	bytes, err := os.ReadFile(path)

	if err != nil {
		logger.Fatal().
			Err(err).
			Str("path", publicKeyPath).
			Msg("Failed to read the public key")
	}

	publicKey := base64.RawURLEncoding.EncodeToString(bytes)

	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		var req sdk.NewClientRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Message: "Invalid JSON body",
			})
		}

		s, err := serverRepository.FindByPrivateIP(ctx, req.Ip)

		if err != nil {
			if err == db.ErrNotFound {
				s, err = serverRepository.Create(ctx, req)
			} else {
				return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
					Message: "Error while creating server",
				})
			}
		}

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Message: "Error while creating server",
			})
		}

		users, err := userRepository.FindByGroup(ctx, s.GroupID)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Message: "Error while creating server",
			})
		}

		userMap := make([]dto.CreateUser, 0, len(users))

		for _, user := range users {
			userMap = append(userMap, dto.CreateUser{
				User: dto.User{
					Name:         user.Name,
					Surname:      user.Surname,
					Username:     user.Username,
					Email:        user.Email,
					Password:     user.Password,
					Shell:        user.Shell,
					SystemGroups: []string{"sudo"},
				},
				PublicSSHKey: user.PublicSSHKey,
			})
		}

		return c.Status(http.StatusCreated).JSON(sdk.NewClientResponse{
			Id:        s.ID,
			PublicKey: publicKey,
		})
	}
}

func DeleteClient(serverRepository server.Interface) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		if err != nil {
			return ctx.Status(http.StatusBadRequest).
				JSON(ErrorResponse{Message: "Invalid Server ID"})
		}

		err = serverRepository.Delete(ctx.UserContext(), uint64(id))

		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Message: "Error while creating server",
			})
		}

		return ctx.SendStatus(http.StatusNoContent)
	}
}
