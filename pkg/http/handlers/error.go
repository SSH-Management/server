package handlers

import (
	"errors"
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/services/password"
)

func Error(logger *log.Logger, translator ut.Translator) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		code := fiber.StatusInternalServerError

		logger.Error().
			Err(err).
			Msg("An error has occurred in application")

		if e, ok := err.(*fiber.Error); ok {
			return ctx.Status(e.Code).JSON(ErrorResponse{
				Message: e.Message,
			})
		}

		if err == ErrInvalidPayload {
			return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Message: ErrInvalidPayload.Error(),
			})
		}

		if err == password.ErrPasswordMismatch {
			return ctx.Status(fiber.StatusUnauthorized).
				JSON(ErrorResponse{Message: "Invalid credentials"})
		}

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).
				JSON(ErrorResponse{Message: "Data is invalid"})
		}

		if err, ok := err.(validator.ValidationErrors); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).
				JSON(err.Translate(translator))
		}

		if errors.Is(err, db.ErrNotFound) {
			return ctx.Status(fiber.StatusNotFound).
				JSON(ErrorResponse{Message: "Data not found!"})
		}

		return ctx.Status(code).
			JSON(ErrorResponse{Message: "An error has occurred!"})
	}
}
