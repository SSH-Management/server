package groups

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/repositories/group"
)

func GetGroups(groupRepo group.Interface) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		groupType := ctx.Query("type")

		var systemGroups bool

		switch groupType {
		case "system":
			systemGroups = true
		case "server":
			systemGroups = false
		default:
			return fiber.NewError(fiber.StatusUnprocessableEntity, "Type not valid for Groups: system or user")
		}

		groups, err := groupRepo.Find(ctx.UserContext(), systemGroups)

		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(groups)
	}
}
