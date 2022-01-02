package handlers

import (
	"errors"
	"github.com/SSH-Management/server/pkg/services/password"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"github.com/SSH-Management/server/pkg/db"
)

type message struct {
	Message string `json:"message"`
}

func Error(logger zerolog.Logger, translator ut.Translator) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		code := fiber.StatusInternalServerError

		logger.Error().
			Err(err).
			Msg("An error has occurred in application")

		if e, ok := err.(*fiber.Error); ok {
			return ctx.Status(e.Code).JSON(message{
				Message: e.Message,
			})
		}

		if err == password.ErrPasswordMismatch {
			return ctx.Status(fiber.StatusUnauthorized).
				JSON(message{Message: "Invalid credentials"})
		}

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).
				JSON(message{Message: "Data is invalid"})
		}

		if err, ok := err.(validator.ValidationErrors); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).
				JSON(err.Translate(translator))
		}

		if errors.Is(err, db.ErrNotFound) {
			return ctx.Status(fiber.StatusNotFound).
				JSON(message{Message: "Data not found!"})
		}

		return ctx.Status(code).
			JSON(message{Message: "An error has occurred!"})
	}
}
